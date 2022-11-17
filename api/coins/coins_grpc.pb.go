// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/coins.proto

package coins

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

// CoinClient is the client API for Coin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinClient interface {
	GetAvailableCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CoinCodes, error)
	SetCoin(ctx context.Context, in *CoinCode, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMyCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CoinCodes, error)
}

type coinClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinClient(cc grpc.ClientConnInterface) CoinClient {
	return &coinClient{cc}
}

func (c *coinClient) GetAvailableCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CoinCodes, error) {
	out := new(CoinCodes)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.coins.coin/GetAvailableCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) SetCoin(ctx context.Context, in *CoinCode, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.coins.coin/SetCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinClient) GetMyCoins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CoinCodes, error) {
	out := new(CoinCodes)
	err := c.cc.Invoke(ctx, "/binance_converter.backend_api.coins.coin/GetMyCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinServer is the server API for Coin service.
// All implementations must embed UnimplementedCoinServer
// for forward compatibility
type CoinServer interface {
	GetAvailableCoins(context.Context, *emptypb.Empty) (*CoinCodes, error)
	SetCoin(context.Context, *CoinCode) (*emptypb.Empty, error)
	GetMyCoins(context.Context, *emptypb.Empty) (*CoinCodes, error)
	mustEmbedUnimplementedCoinServer()
}

// UnimplementedCoinServer must be embedded to have forward compatible implementations.
type UnimplementedCoinServer struct {
}

func (UnimplementedCoinServer) GetAvailableCoins(context.Context, *emptypb.Empty) (*CoinCodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableCoins not implemented")
}
func (UnimplementedCoinServer) SetCoin(context.Context, *CoinCode) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCoin not implemented")
}
func (UnimplementedCoinServer) GetMyCoins(context.Context, *emptypb.Empty) (*CoinCodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCoins not implemented")
}
func (UnimplementedCoinServer) mustEmbedUnimplementedCoinServer() {}

// UnsafeCoinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinServer will
// result in compilation errors.
type UnsafeCoinServer interface {
	mustEmbedUnimplementedCoinServer()
}

func RegisterCoinServer(s grpc.ServiceRegistrar, srv CoinServer) {
	s.RegisterService(&Coin_ServiceDesc, srv)
}

func _Coin_GetAvailableCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).GetAvailableCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.coins.coin/GetAvailableCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).GetAvailableCoins(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coin_SetCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).SetCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.coins.coin/SetCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).SetCoin(ctx, req.(*CoinCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coin_GetMyCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinServer).GetMyCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binance_converter.backend_api.coins.coin/GetMyCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinServer).GetMyCoins(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Coin_ServiceDesc is the grpc.ServiceDesc for Coin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "binance_converter.backend_api.coins.coin",
	HandlerType: (*CoinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAvailableCoins",
			Handler:    _Coin_GetAvailableCoins_Handler,
		},
		{
			MethodName: "SetCoin",
			Handler:    _Coin_SetCoin_Handler,
		},
		{
			MethodName: "GetMyCoins",
			Handler:    _Coin_GetMyCoins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/coins.proto",
}
