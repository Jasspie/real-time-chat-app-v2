// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: proto/chat/v1/chat.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ChatService_SendMsgs_FullMethodName     = "/chat.v1.ChatService/SendMsgs"
	ChatService_CreateUser_FullMethodName   = "/chat.v1.ChatService/CreateUser"
	ChatService_CreateRoom_FullMethodName   = "/chat.v1.ChatService/CreateRoom"
	ChatService_GetRoomUsers_FullMethodName = "/chat.v1.ChatService/GetRoomUsers"
	ChatService_JoinRoom_FullMethodName     = "/chat.v1.ChatService/JoinRoom"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SendMsgs(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMsgsClient, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	GetRoomUsers(ctx context.Context, in *GetRoomUsersRequest, opts ...grpc.CallOption) (ChatService_GetRoomUsersClient, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMsgs(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMsgsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_SendMsgs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMsgsClient{ClientStream: stream}
	return x, nil
}

type ChatService_SendMsgsClient interface {
	Send(*SendMsgsRequest) error
	Recv() (*SendMsgsResponse, error)
	grpc.ClientStream
}

type chatServiceSendMsgsClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMsgsClient) Send(m *SendMsgsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMsgsClient) Recv() (*SendMsgsResponse, error) {
	m := new(SendMsgsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetRoomUsers(ctx context.Context, in *GetRoomUsersRequest, opts ...grpc.CallOption) (ChatService_GetRoomUsersClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_GetRoomUsers_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceGetRoomUsersClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_GetRoomUsersClient interface {
	Recv() (*GetRoomUsersResponse, error)
	grpc.ClientStream
}

type chatServiceGetRoomUsersClient struct {
	grpc.ClientStream
}

func (x *chatServiceGetRoomUsersClient) Recv() (*GetRoomUsersResponse, error) {
	m := new(GetRoomUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinRoomResponse)
	err := c.cc.Invoke(ctx, ChatService_JoinRoom_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SendMsgs(ChatService_SendMsgsServer) error
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	GetRoomUsers(*GetRoomUsersRequest, ChatService_GetRoomUsersServer) error
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SendMsgs(ChatService_SendMsgsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMsgs not implemented")
}
func (UnimplementedChatServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedChatServiceServer) GetRoomUsers(*GetRoomUsersRequest, ChatService_GetRoomUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomUsers not implemented")
}
func (UnimplementedChatServiceServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendMsgs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMsgs(&chatServiceSendMsgsServer{ServerStream: stream})
}

type ChatService_SendMsgsServer interface {
	Send(*SendMsgsResponse) error
	Recv() (*SendMsgsRequest, error)
	grpc.ServerStream
}

type chatServiceSendMsgsServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMsgsServer) Send(m *SendMsgsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMsgsServer) Recv() (*SendMsgsRequest, error) {
	m := new(SendMsgsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetRoomUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRoomUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).GetRoomUsers(m, &chatServiceGetRoomUsersServer{ServerStream: stream})
}

type ChatService_GetRoomUsersServer interface {
	Send(*GetRoomUsersResponse) error
	grpc.ServerStream
}

type chatServiceGetRoomUsersServer struct {
	grpc.ServerStream
}

func (x *chatServiceGetRoomUsersServer) Send(m *GetRoomUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_JoinRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ChatService_CreateUser_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _ChatService_CreateRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _ChatService_JoinRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMsgs",
			Handler:       _ChatService_SendMsgs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetRoomUsers",
			Handler:       _ChatService_GetRoomUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/chat/v1/chat.proto",
}
