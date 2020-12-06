package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jgvkmea/grpc-entry/pb"
	"google.golang.org/grpc"
)

// HelloServer serverの構造体
type HelloServer struct {
}

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	// サーバー登録
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &HelloServer{})
	s.Serve(listenPort)
}

// HelloPerson HelloPersonのメソッド
func (h *HelloServer) HelloPerson(ctx context.Context, req *pb.HelloPersonMessageRequest) (res *pb.HelloPersonMessageResponse, e error) {
	text := fmt.Sprintf("%sさん(%d歳)、こんにちは。", req.Name, req.Age)
	res = &pb.HelloPersonMessageResponse{HelloMessage: text}
	return res, nil
}
