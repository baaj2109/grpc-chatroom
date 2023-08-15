package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "github.com/baaj2109/grpc-chatroom/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedChatRoomServer
	chatMessageCache []*pb.ChatMessage
	userMap          sync.Map
	L                sync.RWMutex
}

var (
	workers map[pb.ChatRoom_ChatServer]pb.ChatRoom_ChatServer = make(map[pb.ChatRoom_ChatServer]pb.ChatRoom_ChatServer)
)

func (s *service) Login(ctx context.Context, user *pb.User) (*wrappers.StringValue, error) {
	user.Id = uuid.New().String()
	if _, ok := s.userMap.Load(user.Id); ok {
		return nil, status.Errorf(codes.AlreadyExists, "已有同名用戶,請換個用戶名")
	}
	s.userMap.Store(user.Id, user)
	go s.sendMessage(nil,
		&pb.ChatMessage{
			Id:      "server",
			Content: fmt.Sprintf("%v 加入聊天室", user.Name),
			Time:    uint64(time.Now().Unix()),
		})
	// some work...
	return &wrappers.StringValue{Value: user.Id}, status.New(codes.OK, "").Err()
}

func (s *service) Chat(stream pb.ChatRoom_ChatServer) error {
	if s.chatMessageCache == nil {
		s.chatMessageCache = make([]*pb.ChatMessage, 0, 1024)
	}
	workers[stream] = stream
	for _, v := range s.chatMessageCache {
		stream.Send(v)
	}
	s.recvMessage(stream)
	return status.New(codes.OK, "").Err()
}

func (s *service) sendMessage(msgSendingServer pb.ChatRoom_ChatServer, mes *pb.ChatMessage) {
	s.L.Lock()
	for _, v := range workers {
		if v != msgSendingServer {
			err := v.Send(mes)
			if err != nil {
				// err handle
				continue
			}
		}
	}
	s.L.Unlock()
}

func (s *service) recvMessage(stream pb.ChatRoom_ChatServer) {
	md, _ := metadata.FromIncomingContext(stream.Context())
	for {
		mes, err := stream.Recv()
		if err != nil {
			if v, ok := s.userMap.Load(md.Get("uuid")[0]); ok {
				s.sendMessage(stream,
					&pb.ChatMessage{
						Id:      "exit",
						Content: fmt.Sprintf("%v 離開聊天室", v.(*pb.User).Name),
						Time:    uint64(time.Now().Unix()),
					})
			}
			s.L.Lock()
			delete(workers, stream)
			s.L.Unlock()
			s.userMap.Delete(md.Get("uuid")[0])
			fmt.Println("用戶離線,目前用戶在線數量", len(workers))
			break
		}
		s.chatMessageCache = append(s.chatMessageCache, mes)
		v, ok := s.userMap.Load(md.Get("uuid")[0])
		if !ok {
			fmt.Println("致命錯誤,用戶不存在")
			return
		}
		mes.Name = v.(*pb.User).Name
		mes.Time = uint64(time.Now().Unix())
		s.sendMessage(stream, mes)
	}
}
