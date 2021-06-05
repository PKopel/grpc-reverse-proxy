package main

import (
	"context"
	"log"
	"net"

	pb "rev_proxy/gen"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) ExampleCall(ctx context.Context, in *pb.ExampleRequest) (*pb.ExampleReply, error) {
	log.Printf("Received: %v, id: %d", in.GetName(), in.GetId())
	return &pb.ExampleReply{Message: "Hello " + in.GetName(), Id: in.GetId()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
