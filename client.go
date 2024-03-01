package main

import (
	"log"

	"github.com/bertoxic/checkgrpc/chat"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
)


func main() {

    var conn *grpc.ClientConn

    conn, err := grpc.Dial(":9000", grpc.WithInsecure(),grpc.WithBlock())
    if err != nil {
        log.Fatal("couldn't connet :%S",err)
    }

    defer conn.Close()  

    c:= chat.NewChatServiceClient(conn)

    message:= chat.Request{
        Request: "Hi i am a client",
        Who: "mark bellai",
    }

    response, err := c.SayHello(context.Background(), &message)
    if err != nil {
        log.Fatal("Error from client :%s",err)
    }

    log.Printf("response from server is %+v\n, from: %s", response.Response, response.Who)

    chatsrv, err := c.Subhello(context.Background(), &message)

    for {
        resp, err := chatsrv.Recv() 
        if err != nil {
            log.Fatalf("Error receiving stream: %v", err)
        }

        log.Printf("Received response: %s\n", resp.Response) 
    }


}