package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jgvkmea/grpc-entry/pb"
	"google.golang.org/grpc"
)

func main() {
	// コネクション作成
	conn, err := grpc.Dial("localhost:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)

	// メソッド呼び出し
	req := &pb.HelloPersonMessageRequest{
		Name: "Qiita",
		Age:  9,
	}
	res, err := client.HelloPerson(context.TODO(), req)
	fmt.Println(res.HelloMessage)
}
