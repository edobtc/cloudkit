// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ResourcesClient is the client API for Resources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourcesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Versions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*VersionsResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CurrentContext(ctx context.Context, in *CurrentContextRequest, opts ...grpc.CallOption) (*CurrentContextResponse, error)
	SetContext(ctx context.Context, in *SetContextRequest, opts ...grpc.CallOption) (*SetContextResponse, error)
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	ProvisionCallback(ctx context.Context, in *ProvisionCallbackRequest, opts ...grpc.CallOption) (*ProvisionCallbackResponse, error)
}

type resourcesClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcesClient(cc grpc.ClientConnInterface) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Versions(ctx context.Context, in *VersionsRequest, opts ...grpc.CallOption) (*VersionsResponse, error) {
	out := new(VersionsResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/Versions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) CurrentContext(ctx context.Context, in *CurrentContextRequest, opts ...grpc.CallOption) (*CurrentContextResponse, error) {
	out := new(CurrentContextResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/CurrentContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) SetContext(ctx context.Context, in *SetContextRequest, opts ...grpc.CallOption) (*SetContextResponse, error) {
	out := new(SetContextResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/SetContext", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) ProvisionCallback(ctx context.Context, in *ProvisionCallbackRequest, opts ...grpc.CallOption) (*ProvisionCallbackResponse, error) {
	out := new(ProvisionCallbackResponse)
	err := c.cc.Invoke(ctx, "/controlplane.resources.v1.Resources/ProvisionCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcesServer is the server API for Resources service.
// All implementations must embed UnimplementedResourcesServer
// for forward compatibility
type ResourcesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Versions(context.Context, *VersionsRequest) (*VersionsResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	CurrentContext(context.Context, *CurrentContextRequest) (*CurrentContextResponse, error)
	SetContext(context.Context, *SetContextRequest) (*SetContextResponse, error)
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	ProvisionCallback(context.Context, *ProvisionCallbackRequest) (*ProvisionCallbackResponse, error)
	mustEmbedUnimplementedResourcesServer()
}

// UnimplementedResourcesServer must be embedded to have forward compatible implementations.
type UnimplementedResourcesServer struct {
}

func (UnimplementedResourcesServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourcesServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedResourcesServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourcesServer) Versions(context.Context, *VersionsRequest) (*VersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Versions not implemented")
}
func (UnimplementedResourcesServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedResourcesServer) CurrentContext(context.Context, *CurrentContextRequest) (*CurrentContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentContext not implemented")
}
func (UnimplementedResourcesServer) SetContext(context.Context, *SetContextRequest) (*SetContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetContext not implemented")
}
func (UnimplementedResourcesServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedResourcesServer) ProvisionCallback(context.Context, *ProvisionCallbackRequest) (*ProvisionCallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionCallback not implemented")
}
func (UnimplementedResourcesServer) mustEmbedUnimplementedResourcesServer() {}

// UnsafeResourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourcesServer will
// result in compilation errors.
type UnsafeResourcesServer interface {
	mustEmbedUnimplementedResourcesServer()
}

func RegisterResourcesServer(s grpc.ServiceRegistrar, srv ResourcesServer) {
	s.RegisterService(&Resources_ServiceDesc, srv)
}

func _Resources_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Versions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Versions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/Versions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Versions(ctx, req.(*VersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_CurrentContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).CurrentContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/CurrentContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).CurrentContext(ctx, req.(*CurrentContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_SetContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).SetContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/SetContext",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).SetContext(ctx, req.(*SetContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_ProvisionCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).ProvisionCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controlplane.resources.v1.Resources/ProvisionCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).ProvisionCallback(ctx, req.(*ProvisionCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Resources_ServiceDesc is the grpc.ServiceDesc for Resources service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resources_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.resources.v1.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Resources_Create_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _Resources_Submit_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Resources_List_Handler,
		},
		{
			MethodName: "Versions",
			Handler:    _Resources_Versions_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Resources_Status_Handler,
		},
		{
			MethodName: "CurrentContext",
			Handler:    _Resources_CurrentContext_Handler,
		},
		{
			MethodName: "SetContext",
			Handler:    _Resources_SetContext_Handler,
		},
		{
			MethodName: "Liveness",
			Handler:    _Resources_Liveness_Handler,
		},
		{
			MethodName: "ProvisionCallback",
			Handler:    _Resources_ProvisionCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/resources/v1/resources.proto",
}
