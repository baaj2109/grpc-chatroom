package main

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	pb "github.com/baaj2109/grpc-chatroom/proto"
	"github.com/pterm/pterm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:23333"
)

func main() {
	/* ---------------------------------- 連接服務器 --------------------------------- */
	spinner, _ := pterm.DefaultSpinner.Start("正在連接聊天室")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		spinner.Fail("連接失敗")
		pterm.Fatal.Printfln("无法連接至服務器: %v", err)
		return
	}
	c := pb.NewChatRoomClient(conn)
	spinner.Success("連接成功")
	/* ---------------------------------- 注册用戶名 --------------------------------- */
	var val *wrapperspb.StringValue
	var user *pb.User
	for {
		result, _ := pterm.DefaultInteractiveTextInput.Show("創建用戶名")
		if strings.TrimSpace(result) == "" {
			pterm.Error.Printfln("進入聊天室失敗,没有取名字")
			continue
		}
		user = &pb.User{Name: result}
		val, err = c.Login(context.TODO(), user)
		if err != nil {
			pterm.Error.Printfln("進入聊天室失敗 %v", err)
			continue
		} else {
			break
		}
	}
	user.Id = val.Value
	pterm.Success.Println("創建成功！開始聊天吧！")
	/* ---------------------------------- 聊天室逻辑 --------------------------------- */
	stream, _ := c.Chat(metadata.AppendToOutgoingContext(context.Background(), "uuid", user.Id))
	go func(pb.ChatRoom_ChatClient) {
		for {
			res, _ := stream.Recv()
			switch res.Id {
			case "server":
				pterm.Success.Printfln("(%[2]v) [server] %[1]s ", res.Content, time.Unix(int64(res.Time), 0).Format(time.ANSIC))
			case "exit":
				pterm.Warning.Printfln("(%[2]v) [User Exit] %[1]s ", res.Content, time.Unix(int64(res.Time), 0).Format(time.ANSIC))

			default:
				pterm.Info.Printfln("(%[3]v) %[1]s : %[2]s", res.Name, res.Content, time.Unix(int64(res.Time), 0).Format(time.ANSIC))
			}
		}
	}(stream)
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimRight(input, "\r \n")
		if input == "exit" {
			// stream.Send(&pb.ChatMessage{Id: user.Id, Content: input})
			break
		}
		// pterm.Info.Printfln("%s : %s", user.Name, input)
		stream.Send(&pb.ChatMessage{Id: user.Id, Content: input})
	}
}
