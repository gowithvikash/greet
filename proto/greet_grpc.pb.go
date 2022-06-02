// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: greet.proto

package proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Simple_Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	Greet_Many_Times(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_Greet_Many_TimesClient, error)
	Long_Greet(ctx context.Context, opts ...grpc.CallOption) (GreetService_Long_GreetClient, error)
	Greet_Every_One(ctx context.Context, opts ...grpc.CallOption) (GreetService_Greet_Every_OneClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Simple_Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Simple_Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) Greet_Many_Times(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_Greet_Many_TimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/Greet_Many_Times", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreet_Many_TimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_Greet_Many_TimesClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreet_Many_TimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreet_Many_TimesClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) Long_Greet(ctx context.Context, opts ...grpc.CallOption) (GreetService_Long_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet.GreetService/Long_Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLong_GreetClient{stream}
	return x, nil
}

type GreetService_Long_GreetClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceLong_GreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLong_GreetClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLong_GreetClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) Greet_Every_One(ctx context.Context, opts ...grpc.CallOption) (GreetService_Greet_Every_OneClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet.GreetService/Greet_Every_One", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreet_Every_OneClient{stream}
	return x, nil
}

type GreetService_Greet_Every_OneClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreet_Every_OneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreet_Every_OneClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreet_Every_OneClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Simple_Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	Greet_Many_Times(*GreetRequest, GreetService_Greet_Many_TimesServer) error
	Long_Greet(GreetService_Long_GreetServer) error
	Greet_Every_One(GreetService_Greet_Every_OneServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Simple_Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simple_Greet not implemented")
}
func (UnimplementedGreetServiceServer) Greet_Many_Times(*GreetRequest, GreetService_Greet_Many_TimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet_Many_Times not implemented")
}
func (UnimplementedGreetServiceServer) Long_Greet(GreetService_Long_GreetServer) error {
	return status.Errorf(codes.Unimplemented, "method Long_Greet not implemented")
}
func (UnimplementedGreetServiceServer) Greet_Every_One(GreetService_Greet_Every_OneServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet_Every_One not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Simple_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Simple_Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Simple_Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Simple_Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_Greet_Many_Times_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).Greet_Many_Times(m, &greetServiceGreet_Many_TimesServer{stream})
}

type GreetService_Greet_Many_TimesServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greetServiceGreet_Many_TimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreet_Many_TimesServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_Long_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).Long_Greet(&greetServiceLong_GreetServer{stream})
}

type GreetService_Long_GreetServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceLong_GreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLong_GreetServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLong_GreetServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_Greet_Every_One_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).Greet_Every_One(&greetServiceGreet_Every_OneServer{stream})
}

type GreetService_Greet_Every_OneServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGreet_Every_OneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreet_Every_OneServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreet_Every_OneServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simple_Greet",
			Handler:    _GreetService_Simple_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet_Many_Times",
			Handler:       _GreetService_Greet_Many_Times_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Long_Greet",
			Handler:       _GreetService_Long_Greet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Greet_Every_One",
			Handler:       _GreetService_Greet_Every_One_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
