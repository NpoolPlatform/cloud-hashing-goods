// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CloudHashingGoodsClient is the client API for CloudHashingGoods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudHashingGoodsClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	CreateVendorLocation(ctx context.Context, in *CreateVendorLocationRequest, opts ...grpc.CallOption) (*CreateVendorLocationResponse, error)
	UpdateVendorLocation(ctx context.Context, in *UpdateVendorLocationRequest, opts ...grpc.CallOption) (*UpdateVendorLocationResponse, error)
	GetVendorLocations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVendorLocationsResponse, error)
	CreateTargetArea(ctx context.Context, in *CreateTargetAreaRequest, opts ...grpc.CallOption) (*CreateTargetAreaResponse, error)
	UpdateTargetArea(ctx context.Context, in *UpdateTargetAreaRequest, opts ...grpc.CallOption) (*UpdateTargetAreaResponse, error)
	GetTargetAreas(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTargetAreasResponse, error)
}

type cloudHashingGoodsClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudHashingGoodsClient(cc grpc.ClientConnInterface) CloudHashingGoodsClient {
	return &cloudHashingGoodsClient{cc}
}

func (c *cloudHashingGoodsClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) CreateVendorLocation(ctx context.Context, in *CreateVendorLocationRequest, opts ...grpc.CallOption) (*CreateVendorLocationResponse, error) {
	out := new(CreateVendorLocationResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateVendorLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) UpdateVendorLocation(ctx context.Context, in *UpdateVendorLocationRequest, opts ...grpc.CallOption) (*UpdateVendorLocationResponse, error) {
	out := new(UpdateVendorLocationResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateVendorLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) GetVendorLocations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVendorLocationsResponse, error) {
	out := new(GetVendorLocationsResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/GetVendorLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) CreateTargetArea(ctx context.Context, in *CreateTargetAreaRequest, opts ...grpc.CallOption) (*CreateTargetAreaResponse, error) {
	out := new(CreateTargetAreaResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/CreateTargetArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) UpdateTargetArea(ctx context.Context, in *UpdateTargetAreaRequest, opts ...grpc.CallOption) (*UpdateTargetAreaResponse, error) {
	out := new(UpdateTargetAreaResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateTargetArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudHashingGoodsClient) GetTargetAreas(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTargetAreasResponse, error) {
	out := new(GetTargetAreasResponse)
	err := c.cc.Invoke(ctx, "/cloud.hashing.goods.v1.CloudHashingGoods/GetTargetAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudHashingGoodsServer is the server API for CloudHashingGoods service.
// All implementations must embed UnimplementedCloudHashingGoodsServer
// for forward compatibility
type CloudHashingGoodsServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	CreateVendorLocation(context.Context, *CreateVendorLocationRequest) (*CreateVendorLocationResponse, error)
	UpdateVendorLocation(context.Context, *UpdateVendorLocationRequest) (*UpdateVendorLocationResponse, error)
	GetVendorLocations(context.Context, *emptypb.Empty) (*GetVendorLocationsResponse, error)
	CreateTargetArea(context.Context, *CreateTargetAreaRequest) (*CreateTargetAreaResponse, error)
	UpdateTargetArea(context.Context, *UpdateTargetAreaRequest) (*UpdateTargetAreaResponse, error)
	GetTargetAreas(context.Context, *emptypb.Empty) (*GetTargetAreasResponse, error)
	mustEmbedUnimplementedCloudHashingGoodsServer()
}

// UnimplementedCloudHashingGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedCloudHashingGoodsServer struct {
}

func (UnimplementedCloudHashingGoodsServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCloudHashingGoodsServer) CreateVendorLocation(context.Context, *CreateVendorLocationRequest) (*CreateVendorLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVendorLocation not implemented")
}
func (UnimplementedCloudHashingGoodsServer) UpdateVendorLocation(context.Context, *UpdateVendorLocationRequest) (*UpdateVendorLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVendorLocation not implemented")
}
func (UnimplementedCloudHashingGoodsServer) GetVendorLocations(context.Context, *emptypb.Empty) (*GetVendorLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVendorLocations not implemented")
}
func (UnimplementedCloudHashingGoodsServer) CreateTargetArea(context.Context, *CreateTargetAreaRequest) (*CreateTargetAreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTargetArea not implemented")
}
func (UnimplementedCloudHashingGoodsServer) UpdateTargetArea(context.Context, *UpdateTargetAreaRequest) (*UpdateTargetAreaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTargetArea not implemented")
}
func (UnimplementedCloudHashingGoodsServer) GetTargetAreas(context.Context, *emptypb.Empty) (*GetTargetAreasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargetAreas not implemented")
}
func (UnimplementedCloudHashingGoodsServer) mustEmbedUnimplementedCloudHashingGoodsServer() {}

// UnsafeCloudHashingGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudHashingGoodsServer will
// result in compilation errors.
type UnsafeCloudHashingGoodsServer interface {
	mustEmbedUnimplementedCloudHashingGoodsServer()
}

func RegisterCloudHashingGoodsServer(s grpc.ServiceRegistrar, srv CloudHashingGoodsServer) {
	s.RegisterService(&CloudHashingGoods_ServiceDesc, srv)
}

func _CloudHashingGoods_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_CreateVendorLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVendorLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).CreateVendorLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/CreateVendorLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).CreateVendorLocation(ctx, req.(*CreateVendorLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_UpdateVendorLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVendorLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).UpdateVendorLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateVendorLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).UpdateVendorLocation(ctx, req.(*UpdateVendorLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_GetVendorLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).GetVendorLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/GetVendorLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).GetVendorLocations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_CreateTargetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTargetAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).CreateTargetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/CreateTargetArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).CreateTargetArea(ctx, req.(*CreateTargetAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_UpdateTargetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTargetAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).UpdateTargetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/UpdateTargetArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).UpdateTargetArea(ctx, req.(*UpdateTargetAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudHashingGoods_GetTargetAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudHashingGoodsServer).GetTargetAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.hashing.goods.v1.CloudHashingGoods/GetTargetAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudHashingGoodsServer).GetTargetAreas(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudHashingGoods_ServiceDesc is the grpc.ServiceDesc for CloudHashingGoods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudHashingGoods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.hashing.goods.v1.CloudHashingGoods",
	HandlerType: (*CloudHashingGoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _CloudHashingGoods_Version_Handler,
		},
		{
			MethodName: "CreateVendorLocation",
			Handler:    _CloudHashingGoods_CreateVendorLocation_Handler,
		},
		{
			MethodName: "UpdateVendorLocation",
			Handler:    _CloudHashingGoods_UpdateVendorLocation_Handler,
		},
		{
			MethodName: "GetVendorLocations",
			Handler:    _CloudHashingGoods_GetVendorLocations_Handler,
		},
		{
			MethodName: "CreateTargetArea",
			Handler:    _CloudHashingGoods_CreateTargetArea_Handler,
		},
		{
			MethodName: "UpdateTargetArea",
			Handler:    _CloudHashingGoods_UpdateTargetArea_Handler,
		},
		{
			MethodName: "GetTargetAreas",
			Handler:    _CloudHashingGoods_GetTargetAreas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/cloud-hashing-goods.proto",
}
