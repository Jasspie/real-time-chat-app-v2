// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: chat/v1/chat.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/Jasspie/real-time-chat-app-v2/server/chat/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ChatServiceName is the fully-qualified name of the ChatService service.
	ChatServiceName = "chat.v1.ChatService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChatServiceNewChatSessionProcedure is the fully-qualified name of the ChatService's
	// NewChatSession RPC.
	ChatServiceNewChatSessionProcedure = "/chat.v1.ChatService/NewChatSession"
	// ChatServiceBroadcastChatProcedure is the fully-qualified name of the ChatService's BroadcastChat
	// RPC.
	ChatServiceBroadcastChatProcedure = "/chat.v1.ChatService/BroadcastChat"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	chatServiceServiceDescriptor              = v1.File_chat_v1_chat_proto.Services().ByName("ChatService")
	chatServiceNewChatSessionMethodDescriptor = chatServiceServiceDescriptor.Methods().ByName("NewChatSession")
	chatServiceBroadcastChatMethodDescriptor  = chatServiceServiceDescriptor.Methods().ByName("BroadcastChat")
)

// ChatServiceClient is a client for the chat.v1.ChatService service.
type ChatServiceClient interface {
	NewChatSession(context.Context, *connect.Request[v1.NewChatSessionRequest]) (*connect.ServerStreamForClient[v1.NewChatSessionResponse], error)
	BroadcastChat(context.Context, *connect.Request[v1.BroadcastChatRequest]) (*connect.Response[v1.BroadcastChatResponse], error)
}

// NewChatServiceClient constructs a client for the chat.v1.ChatService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChatServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChatServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &chatServiceClient{
		newChatSession: connect.NewClient[v1.NewChatSessionRequest, v1.NewChatSessionResponse](
			httpClient,
			baseURL+ChatServiceNewChatSessionProcedure,
			connect.WithSchema(chatServiceNewChatSessionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		broadcastChat: connect.NewClient[v1.BroadcastChatRequest, v1.BroadcastChatResponse](
			httpClient,
			baseURL+ChatServiceBroadcastChatProcedure,
			connect.WithSchema(chatServiceBroadcastChatMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// chatServiceClient implements ChatServiceClient.
type chatServiceClient struct {
	newChatSession *connect.Client[v1.NewChatSessionRequest, v1.NewChatSessionResponse]
	broadcastChat  *connect.Client[v1.BroadcastChatRequest, v1.BroadcastChatResponse]
}

// NewChatSession calls chat.v1.ChatService.NewChatSession.
func (c *chatServiceClient) NewChatSession(ctx context.Context, req *connect.Request[v1.NewChatSessionRequest]) (*connect.ServerStreamForClient[v1.NewChatSessionResponse], error) {
	return c.newChatSession.CallServerStream(ctx, req)
}

// BroadcastChat calls chat.v1.ChatService.BroadcastChat.
func (c *chatServiceClient) BroadcastChat(ctx context.Context, req *connect.Request[v1.BroadcastChatRequest]) (*connect.Response[v1.BroadcastChatResponse], error) {
	return c.broadcastChat.CallUnary(ctx, req)
}

// ChatServiceHandler is an implementation of the chat.v1.ChatService service.
type ChatServiceHandler interface {
	NewChatSession(context.Context, *connect.Request[v1.NewChatSessionRequest], *connect.ServerStream[v1.NewChatSessionResponse]) error
	BroadcastChat(context.Context, *connect.Request[v1.BroadcastChatRequest]) (*connect.Response[v1.BroadcastChatResponse], error)
}

// NewChatServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChatServiceHandler(svc ChatServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	chatServiceNewChatSessionHandler := connect.NewServerStreamHandler(
		ChatServiceNewChatSessionProcedure,
		svc.NewChatSession,
		connect.WithSchema(chatServiceNewChatSessionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	chatServiceBroadcastChatHandler := connect.NewUnaryHandler(
		ChatServiceBroadcastChatProcedure,
		svc.BroadcastChat,
		connect.WithSchema(chatServiceBroadcastChatMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/chat.v1.ChatService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChatServiceNewChatSessionProcedure:
			chatServiceNewChatSessionHandler.ServeHTTP(w, r)
		case ChatServiceBroadcastChatProcedure:
			chatServiceBroadcastChatHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChatServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChatServiceHandler struct{}

func (UnimplementedChatServiceHandler) NewChatSession(context.Context, *connect.Request[v1.NewChatSessionRequest], *connect.ServerStream[v1.NewChatSessionResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.NewChatSession is not implemented"))
}

func (UnimplementedChatServiceHandler) BroadcastChat(context.Context, *connect.Request[v1.BroadcastChatRequest]) (*connect.Response[v1.BroadcastChatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("chat.v1.ChatService.BroadcastChat is not implemented"))
}
