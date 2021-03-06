package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	gen "rev_proxy/gen"

	"google.golang.org/grpc"
)

const (
	defaultPort = ":50051"
	defaultName = "server"
)

type exampleServer struct {
	gen.UnimplementedExampleServiceServer
	name string
}

func (server *exampleServer) ExampleCall(ctx context.Context, in *gen.ExampleRequest) (*gen.ExampleReply, error) {
	log.Printf("Received: %v, id: %d", in.GetName(), in.GetId())
	message := fmt.Sprintf("reply from %v to %v", server.name, in.GetName())
	return &gen.ExampleReply{Message: message, Id: in.GetId()}, nil
}

func main() {
	port := defaultPort
	name := defaultName
	switch len(os.Args) {
	case 3:
		port = os.Args[1]
		name = os.Args[2]
	case 2:
		port = os.Args[1]
	}
	log.Printf("port: %v name: %v", port, name)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	gen.RegisterExampleServiceServer(server, &exampleServer{name: name})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
