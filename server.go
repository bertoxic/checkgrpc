package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bertoxic/checkgrpc/chat"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("in main package now")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("cannot connect to server net", err)
		return
	}
	s := &chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, s)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc %v", err)
	}
}
