//go:generate sh -c "protoc -I $PWD/../../proto --go_out=plugins=grpc:$PWD/helloworld $PWD/../../proto/*.proto"

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/ryutah/grpc-helloworld/server/go/helloworld"
)

type greeterServer struct{}

func (g *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	var name string
	if len(in.GetName()) > 0 {
		name = in.GetName()
	} else {
		name = "no name"
	}

	log.Printf("Receive name: %q", name)
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s!!", name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("faield to listen port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, new(greeterServer))

	reflection.Register(server)

	log.Println("Start server on port 8080\nPlease press Ctrl-C to stop server")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
