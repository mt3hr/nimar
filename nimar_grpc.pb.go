// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: nimar.proto

package nimar

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

// NimaRClient is the client API for NimaR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NimaRClient interface {
	ListRooms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Rooms, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*Room, error)
	GameTableStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_GameTableStreamClient, error)
	GetPlayerID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlayerID, error)
	MessageStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_MessageStreamClient, error)
	OperatorsStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_OperatorsStreamClient, error)
	Operate(ctx context.Context, in *Operator, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type nimaRClient struct {
	cc grpc.ClientConnInterface
}

func NewNimaRClient(cc grpc.ClientConnInterface) NimaRClient {
	return &nimaRClient{cc}
}

func (c *nimaRClient) ListRooms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Rooms, error) {
	out := new(Rooms)
	err := c.cc.Invoke(ctx, "/NimaR/ListRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nimaRClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/NimaR/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nimaRClient) GameTableStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_GameTableStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NimaR_ServiceDesc.Streams[0], "/NimaR/GameTableStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nimaRGameTableStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NimaR_GameTableStreamClient interface {
	Recv() (*GameTable, error)
	grpc.ClientStream
}

type nimaRGameTableStreamClient struct {
	grpc.ClientStream
}

func (x *nimaRGameTableStreamClient) Recv() (*GameTable, error) {
	m := new(GameTable)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nimaRClient) GetPlayerID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlayerID, error) {
	out := new(PlayerID)
	err := c.cc.Invoke(ctx, "/NimaR/GetPlayerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nimaRClient) MessageStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NimaR_ServiceDesc.Streams[1], "/NimaR/MessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nimaRMessageStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NimaR_MessageStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type nimaRMessageStreamClient struct {
	grpc.ClientStream
}

func (x *nimaRMessageStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nimaRClient) OperatorsStream(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (NimaR_OperatorsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NimaR_ServiceDesc.Streams[2], "/NimaR/OperatorsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &nimaROperatorsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NimaR_OperatorsStreamClient interface {
	Recv() (*Operators, error)
	grpc.ClientStream
}

type nimaROperatorsStreamClient struct {
	grpc.ClientStream
}

func (x *nimaROperatorsStreamClient) Recv() (*Operators, error) {
	m := new(Operators)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nimaRClient) Operate(ctx context.Context, in *Operator, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/NimaR/Operate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NimaRServer is the server API for NimaR service.
// All implementations must embed UnimplementedNimaRServer
// for forward compatibility
type NimaRServer interface {
	ListRooms(context.Context, *emptypb.Empty) (*Rooms, error)
	CreateRoom(context.Context, *CreateRoomRequest) (*Room, error)
	GameTableStream(*JoinRoomRequest, NimaR_GameTableStreamServer) error
	GetPlayerID(context.Context, *emptypb.Empty) (*PlayerID, error)
	MessageStream(*JoinRoomRequest, NimaR_MessageStreamServer) error
	OperatorsStream(*JoinRoomRequest, NimaR_OperatorsStreamServer) error
	Operate(context.Context, *Operator) (*emptypb.Empty, error)
	mustEmbedUnimplementedNimaRServer()
}

// UnimplementedNimaRServer must be embedded to have forward compatible implementations.
type UnimplementedNimaRServer struct {
}

func (UnimplementedNimaRServer) ListRooms(context.Context, *emptypb.Empty) (*Rooms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}
func (UnimplementedNimaRServer) CreateRoom(context.Context, *CreateRoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedNimaRServer) GameTableStream(*JoinRoomRequest, NimaR_GameTableStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GameTableStream not implemented")
}
func (UnimplementedNimaRServer) GetPlayerID(context.Context, *emptypb.Empty) (*PlayerID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerID not implemented")
}
func (UnimplementedNimaRServer) MessageStream(*JoinRoomRequest, NimaR_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}
func (UnimplementedNimaRServer) OperatorsStream(*JoinRoomRequest, NimaR_OperatorsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OperatorsStream not implemented")
}
func (UnimplementedNimaRServer) Operate(context.Context, *Operator) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Operate not implemented")
}
func (UnimplementedNimaRServer) mustEmbedUnimplementedNimaRServer() {}

// UnsafeNimaRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NimaRServer will
// result in compilation errors.
type UnsafeNimaRServer interface {
	mustEmbedUnimplementedNimaRServer()
}

func RegisterNimaRServer(s grpc.ServiceRegistrar, srv NimaRServer) {
	s.RegisterService(&NimaR_ServiceDesc, srv)
}

func _NimaR_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NimaRServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NimaR/ListRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NimaRServer).ListRooms(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NimaR_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NimaRServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NimaR/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NimaRServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NimaR_GameTableStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NimaRServer).GameTableStream(m, &nimaRGameTableStreamServer{stream})
}

type NimaR_GameTableStreamServer interface {
	Send(*GameTable) error
	grpc.ServerStream
}

type nimaRGameTableStreamServer struct {
	grpc.ServerStream
}

func (x *nimaRGameTableStreamServer) Send(m *GameTable) error {
	return x.ServerStream.SendMsg(m)
}

func _NimaR_GetPlayerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NimaRServer).GetPlayerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NimaR/GetPlayerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NimaRServer).GetPlayerID(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NimaR_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NimaRServer).MessageStream(m, &nimaRMessageStreamServer{stream})
}

type NimaR_MessageStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type nimaRMessageStreamServer struct {
	grpc.ServerStream
}

func (x *nimaRMessageStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _NimaR_OperatorsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinRoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NimaRServer).OperatorsStream(m, &nimaROperatorsStreamServer{stream})
}

type NimaR_OperatorsStreamServer interface {
	Send(*Operators) error
	grpc.ServerStream
}

type nimaROperatorsStreamServer struct {
	grpc.ServerStream
}

func (x *nimaROperatorsStreamServer) Send(m *Operators) error {
	return x.ServerStream.SendMsg(m)
}

func _NimaR_Operate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NimaRServer).Operate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NimaR/Operate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NimaRServer).Operate(ctx, req.(*Operator))
	}
	return interceptor(ctx, in, info, handler)
}

// NimaR_ServiceDesc is the grpc.ServiceDesc for NimaR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NimaR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NimaR",
	HandlerType: (*NimaRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRooms",
			Handler:    _NimaR_ListRooms_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _NimaR_CreateRoom_Handler,
		},
		{
			MethodName: "GetPlayerID",
			Handler:    _NimaR_GetPlayerID_Handler,
		},
		{
			MethodName: "Operate",
			Handler:    _NimaR_Operate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GameTableStream",
			Handler:       _NimaR_GameTableStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MessageStream",
			Handler:       _NimaR_MessageStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OperatorsStream",
			Handler:       _NimaR_OperatorsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nimar.proto",
}
