package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/uxsnap/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (c *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	getResp := &desc.CreateRequest{
		Usernames: []string{"Yes", "yes1", "yes2"},
	}

	fmt.Println(getResp)

	return &desc.CreateResponse{
		Id: int64(len(getResp.Usernames)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))

	if err != nil {
		log.Fatal("Cannot create tcp connection!")
		return
	}

	grpcS := grpc.NewServer()
	reflection.Register(grpcS)

	if err != nil {
		log.Fatal("Cannot create grpc connection!")
		return
	}

	desc.RegisterChatV1Server(grpcS, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = grpcS.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
