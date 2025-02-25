package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/loft-sh/log"
	"github.com/loft-sh/log/survey"
	"github.com/loft-sh/vcluster/pkg/cli/find"
	"github.com/loft-sh/vcluster/pkg/cli/flags"
	"github.com/loft-sh/vcluster/pkg/lifecycle"
	"github.com/loft-sh/vcluster/pkg/platform"
	"github.com/loft-sh/vcluster/pkg/platform/clihelper"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

type AddVClusterOptions struct {
	Project                  string
	ImportName               string
	Restart                  bool
	Insecure                 bool
	AccessKey                string
	Host                     string
	CertificateAuthorityData []byte
	All                      bool
}

func AddVClusterHelm(
	ctx context.Context,
	options *AddVClusterOptions,
	globalFlags *flags.GlobalFlags,
	args []string,
	log log.Logger,
) error {
	var vClusters []find.VCluster
	if len(args) == 0 && !options.All {
		return errors.New("empty vCluster name but no --all flag set, please either set vCluster name to add one cluster or set --all flag to add all of them")
	}
	if options.All {
		log.Info("looking for vCluster instances in all namespaces")
		vClustersInNamespace, err := find.ListVClusters(ctx, globalFlags.Context, "", "", log)
		if err != nil {
			return err
		}
		if len(vClustersInNamespace) == 0 {
			log.Infof("no vCluster instances found in context %s", globalFlags.Context)
		} else {
			vClusters = append(vClusters, vClustersInNamespace...)
		}
	} else {
		// check if vCluster exists
		vClusterName := args[0]
		vCluster, err := find.GetVCluster(ctx, globalFlags.Context, vClusterName, globalFlags.Namespace, log)
		if err != nil {
			return err
		}
		vClusters = append(vClusters, *vCluster)
	}

	if len(vClusters) == 0 {
		return nil
	}

	restConfig, err := vClusters[0].ClientFactory.ClientConfig()
	if err != nil {
		return err
	}

	// create kube client
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}
	var addErrors []error
	log.Debugf("trying to add %d vCluster instances to platform", len(vClusters))
	for _, vCluster := range vClusters {
		log.Infof("adding %s vCluster to platform", vCluster.Name)
		err := addVClusterHelm(ctx, options, globalFlags, vCluster.Name, &vCluster, kubeClient, log)
		if err != nil {
			addErrors = append(addErrors, fmt.Errorf("cannot add %s: %w", vCluster.Name, err))
		}
	}

	return utilerrors.NewAggregate(addErrors)
}

func addVClusterHelm(
	ctx context.Context,
	options *AddVClusterOptions,
	globalFlags *flags.GlobalFlags,
	vClusterName string,
	vCluster *find.VCluster,
	kubeClient *kubernetes.Clientset,
	log log.Logger,
) error {
	snoozed := false
	// If the vCluster was paused with the helm driver, adding it to the platform will only create the secret for registration
	// which leads to confusing behavior for the user since they won't see the cluster in the platform UI until it is resumed.
	if lifecycle.IsPaused(vCluster) {
		log.Infof("vCluster %s is currently sleeping. It will not be added to the platform until it wakes again.", vCluster.Name)

		snoozeConfirmation := "No. Leave it sleeping. (It will be added automatically on next wakeup)"
		answer, err := log.Question(&survey.QuestionOptions{
			Question:     fmt.Sprintf("Would you like to wake vCluster %s now to add immediately?", vCluster.Name),
			DefaultValue: snoozeConfirmation,
			Options: []string{
				snoozeConfirmation,
				"Yes. Wake and add now.",
			},
		})
		if err != nil {
			return fmt.Errorf("failed to capture your response %w", err)
		}

		if snoozed = answer == snoozeConfirmation; !snoozed {
			if err = ResumeHelm(ctx, globalFlags, vClusterName, log); err != nil {
				return fmt.Errorf("failed to wake up vCluster %s: %w", vClusterName, err)
			}

			err = wait.PollUntilContextTimeout(ctx, time.Second, clihelper.Timeout(), false, func(ctx context.Context) (done bool, err error) {
				vCluster, err = find.GetVCluster(ctx, globalFlags.Context, vClusterName, globalFlags.Namespace, log)
				if err != nil {
					return false, err
				}

				return !lifecycle.IsPaused(vCluster), nil
			})
			if err != nil {
				return fmt.Errorf("error waiting for vCluster to wake up %w", err)
			}
		}
	}

	// apply platform secret
	err := platform.ApplyPlatformSecret(
		ctx,
		globalFlags.LoadedConfig(log),
		kubeClient,
		options.ImportName,
		vCluster.Name,
		vCluster.Namespace,
		options.Project,
		options.AccessKey,
		options.Host,
		options.Insecure,
		options.CertificateAuthorityData,
		log,
	)
	if err != nil {
		return err
	}

	// restart vCluster
	if options.Restart {
		log.Infof("Restarting vCluster")
		err = lifecycle.DeletePods(ctx, kubeClient, "app=vcluster,release="+vCluster.Name, vCluster.Namespace)
		if err != nil {
			return fmt.Errorf("delete vcluster workloads: %w", err)
		}
	}

	if snoozed {
		log.Infof("vCluster %s/%s will be added the next time it awakes", vCluster.Namespace, vCluster.Name)
		log.Donef("Run 'vcluster wakeup --help' to learn how to wake up vCluster %s/%s to complete the add operation.", vCluster.Namespace, vCluster.Name)
	} else {
		log.Donef("Successfully added vCluster %s/%s", vCluster.Namespace, vCluster.Name)
	}
	return nil
}
