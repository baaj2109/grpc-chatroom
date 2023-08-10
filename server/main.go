package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/baaj2109/grpc-chatroom/proto"
	"google.golang.org/grpc"
)

const (
	ip   = "127.0.0.1"
	port = "23333"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		log.Fatalf("无法监听端口 %v %v", port, err)
	}
	s := grpc.NewServer()
	// ^ 注册服务
	pb.RegisterChatRoomServer(s, &service{})
	log.Println("gRPC服务器开始监听", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("提供服务失败: %v", err)
	}
}
