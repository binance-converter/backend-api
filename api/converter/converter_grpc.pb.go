// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/converter.proto

package converter

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

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConverterClient interface {
	GetAvailableConverterPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConverterPairs, error)
	SetConvertPair(ctx context.Context, in *ConverterPair, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMyConvertPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConverterPairs, error)
	SetThresholdConvertPairs(ctx context.Context, in *ThresholdConvertPair, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMyThresholdConvertPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ThresholdConvertPairs, error)
	GetCurrentExchange(ctx context.Context, in *ConverterPair, opts ...grpc.CallOption) (*Exchange, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) GetAvailableConverterPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConverterPairs, error) {
	out := new(ConverterPairs)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/GetAvailableConverterPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) SetConvertPair(ctx context.Context, in *ConverterPair, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/SetConvertPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) GetMyConvertPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConverterPairs, error) {
	out := new(ConverterPairs)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/GetMyConvertPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) SetThresholdConvertPairs(ctx context.Context, in *ThresholdConvertPair, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/SetThresholdConvertPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) GetMyThresholdConvertPairs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ThresholdConvertPairs, error) {
	out := new(ThresholdConvertPairs)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/GetMyThresholdConvertPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) GetCurrentExchange(ctx context.Context, in *ConverterPair, opts ...grpc.CallOption) (*Exchange, error) {
	out := new(Exchange)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.converter.converter/GetCurrentExchange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConverterServer is the server API for Converter service.
// All implementations must embed UnimplementedConverterServer
// for forward compatibility
type ConverterServer interface {
	GetAvailableConverterPairs(context.Context, *emptypb.Empty) (*ConverterPairs, error)
	SetConvertPair(context.Context, *ConverterPair) (*emptypb.Empty, error)
	GetMyConvertPairs(context.Context, *emptypb.Empty) (*ConverterPairs, error)
	SetThresholdConvertPairs(context.Context, *ThresholdConvertPair) (*emptypb.Empty, error)
	GetMyThresholdConvertPairs(context.Context, *emptypb.Empty) (*ThresholdConvertPairs, error)
	GetCurrentExchange(context.Context, *ConverterPair) (*Exchange, error)
	mustEmbedUnimplementedConverterServer()
}

// UnimplementedConverterServer must be embedded to have forward compatible implementations.
type UnimplementedConverterServer struct {
}

func (UnimplementedConverterServer) GetAvailableConverterPairs(context.Context, *emptypb.Empty) (*ConverterPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableConverterPairs not implemented")
}
func (UnimplementedConverterServer) SetConvertPair(context.Context, *ConverterPair) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConvertPair not implemented")
}
func (UnimplementedConverterServer) GetMyConvertPairs(context.Context, *emptypb.Empty) (*ConverterPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyConvertPairs not implemented")
}
func (UnimplementedConverterServer) SetThresholdConvertPairs(context.Context, *ThresholdConvertPair) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetThresholdConvertPairs not implemented")
}
func (UnimplementedConverterServer) GetMyThresholdConvertPairs(context.Context, *emptypb.Empty) (*ThresholdConvertPairs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyThresholdConvertPairs not implemented")
}
func (UnimplementedConverterServer) GetCurrentExchange(context.Context, *ConverterPair) (*Exchange, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentExchange not implemented")
}
func (UnimplementedConverterServer) mustEmbedUnimplementedConverterServer() {}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_GetAvailableConverterPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).GetAvailableConverterPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/GetAvailableConverterPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).GetAvailableConverterPairs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_SetConvertPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConverterPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).SetConvertPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/SetConvertPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).SetConvertPair(ctx, req.(*ConverterPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_GetMyConvertPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).GetMyConvertPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/GetMyConvertPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).GetMyConvertPairs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_SetThresholdConvertPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThresholdConvertPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).SetThresholdConvertPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/SetThresholdConvertPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).SetThresholdConvertPairs(ctx, req.(*ThresholdConvertPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_GetMyThresholdConvertPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).GetMyThresholdConvertPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/GetMyThresholdConvertPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).GetMyThresholdConvertPairs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_GetCurrentExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConverterPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).GetCurrentExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.converter.converter/GetCurrentExchange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).GetCurrentExchange(ctx, req.(*ConverterPair))
	}
	return interceptor(ctx, in, info, handler)
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "binance_converter.backend_api.converter.converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailableConverterPairs",
			Handler:    _Converter_GetAvailableConverterPairs_Handler,
		},
		{
			MethodName: "SetConvertPair",
			Handler:    _Converter_SetConvertPair_Handler,
		},
		{
			MethodName: "GetMyConvertPairs",
			Handler:    _Converter_GetMyConvertPairs_Handler,
		},
		{
			MethodName: "SetThresholdConvertPairs",
			Handler:    _Converter_SetThresholdConvertPairs_Handler,
		},
		{
			MethodName: "GetMyThresholdConvertPairs",
			Handler:    _Converter_GetMyThresholdConvertPairs_Handler,
		},
		{
			MethodName: "GetCurrentExchange",
			Handler:    _Converter_GetCurrentExchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/converter.proto",
}
