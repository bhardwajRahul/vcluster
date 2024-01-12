// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: pluginv2.proto

package pluginv2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	Start(ctx context.Context, in *Start_Request, opts ...grpc.CallOption) (*Start_Response, error)
	GetClientHooks(ctx context.Context, in *GetClientHooks_Request, opts ...grpc.CallOption) (*GetClientHooks_Response, error)
	Mutate(ctx context.Context, in *Mutate_Request, opts ...grpc.CallOption) (*Mutate_Response, error)
	SetLeader(ctx context.Context, in *SetLeader_Request, opts ...grpc.CallOption) (*SetLeader_Response, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Start(ctx context.Context, in *Start_Request, opts ...grpc.CallOption) (*Start_Response, error) {
	out := new(Start_Response)
	err := c.cc.Invoke(ctx, "/pluginv2.Plugin/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetClientHooks(ctx context.Context, in *GetClientHooks_Request, opts ...grpc.CallOption) (*GetClientHooks_Response, error) {
	out := new(GetClientHooks_Response)
	err := c.cc.Invoke(ctx, "/pluginv2.Plugin/GetClientHooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Mutate(ctx context.Context, in *Mutate_Request, opts ...grpc.CallOption) (*Mutate_Response, error) {
	out := new(Mutate_Response)
	err := c.cc.Invoke(ctx, "/pluginv2.Plugin/Mutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) SetLeader(ctx context.Context, in *SetLeader_Request, opts ...grpc.CallOption) (*SetLeader_Response, error) {
	out := new(SetLeader_Response)
	err := c.cc.Invoke(ctx, "/pluginv2.Plugin/SetLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	Start(context.Context, *Start_Request) (*Start_Response, error)
	GetClientHooks(context.Context, *GetClientHooks_Request) (*GetClientHooks_Response, error)
	Mutate(context.Context, *Mutate_Request) (*Mutate_Response, error)
	SetLeader(context.Context, *SetLeader_Request) (*SetLeader_Response, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) Start(context.Context, *Start_Request) (*Start_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedPluginServer) GetClientHooks(context.Context, *GetClientHooks_Request) (*GetClientHooks_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientHooks not implemented")
}
func (UnimplementedPluginServer) Mutate(context.Context, *Mutate_Request) (*Mutate_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate not implemented")
}
func (UnimplementedPluginServer) SetLeader(context.Context, *SetLeader_Request) (*SetLeader_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLeader not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Start_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginv2.Plugin/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Start(ctx, req.(*Start_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetClientHooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientHooks_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetClientHooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginv2.Plugin/GetClientHooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetClientHooks(ctx, req.(*GetClientHooks_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mutate_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginv2.Plugin/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Mutate(ctx, req.(*Mutate_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_SetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLeader_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).SetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginv2.Plugin/SetLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).SetLeader(ctx, req.(*SetLeader_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Plugin_Start_Handler,
		},
		{
			MethodName: "GetClientHooks",
			Handler:    _Plugin_GetClientHooks_Handler,
		},
		{
			MethodName: "Mutate",
			Handler:    _Plugin_Mutate_Handler,
		},
		{
			MethodName: "SetLeader",
			Handler:    _Plugin_SetLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pluginv2.proto",
}