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
		log.Fatalf("無法監聽端口 %v %v", port, err)
	}
	s := grpc.NewServer()
	// 註冊服務
	pb.RegisterChatRoomServer(s, &service{})
	log.Println("gRPC 開始監聽", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("提供服務失敗: %v", err)
	}
}
