// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: chat_room.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatRoom_Login_FullMethodName = "/chatroom.ChatRoom/login"
	ChatRoom_Chat_FullMethodName  = "/chatroom.ChatRoom/chat"
)

// ChatRoomClient is the client API for ChatRoom service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatRoomClient interface {
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error)
}

type chatRoomClient struct {
	cc grpc.ClientConnInterface
}

func NewChatRoomClient(cc grpc.ClientConnInterface) ChatRoomClient {
	return &chatRoomClient{cc}
}

func (c *chatRoomClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, ChatRoom_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatRoomClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatRoom_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatRoom_ServiceDesc.Streams[0], ChatRoom_Chat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatRoomChatClient{stream}
	return x, nil
}

type ChatRoom_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatRoomChatClient struct {
	grpc.ClientStream
}

func (x *chatRoomChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatRoomChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatRoomServer is the server API for ChatRoom service.
// All implementations must embed UnimplementedChatRoomServer
// for forward compatibility
type ChatRoomServer interface {
	Login(context.Context, *User) (*wrapperspb.StringValue, error)
	Chat(ChatRoom_ChatServer) error
	mustEmbedUnimplementedChatRoomServer()
}

// UnimplementedChatRoomServer must be embedded to have forward compatible implementations.
type UnimplementedChatRoomServer struct {
}

func (UnimplementedChatRoomServer) Login(context.Context, *User) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatRoomServer) Chat(ChatRoom_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatRoomServer) mustEmbedUnimplementedChatRoomServer() {}

// UnsafeChatRoomServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatRoomServer will
// result in compilation errors.
type UnsafeChatRoomServer interface {
	mustEmbedUnimplementedChatRoomServer()
}

func RegisterChatRoomServer(s grpc.ServiceRegistrar, srv ChatRoomServer) {
	s.RegisterService(&ChatRoom_ServiceDesc, srv)
}

func _ChatRoom_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatRoomServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatRoom_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatRoomServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatRoom_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatRoomServer).Chat(&chatRoomChatServer{stream})
}

type ChatRoom_ChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatRoomChatServer struct {
	grpc.ServerStream
}

func (x *chatRoomChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatRoomChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatRoom_ServiceDesc is the grpc.ServiceDesc for ChatRoom service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatRoom_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatroom.ChatRoom",
	HandlerType: (*ChatRoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _ChatRoom_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "chat",
			Handler:       _ChatRoom_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat_room.proto",
}
