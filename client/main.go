package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "rev_proxy/gen"

	"google.golang.org/grpc"
)

const (
	defaultAddress = "localhost:50051"
	defaultName    = "client"
)

func main() {

	address := defaultAddress
	name := defaultName
	switch len(os.Args) {
	case 3:
		name = os.Args[1]
		address = os.Args[2]
	case 2:
		name = os.Args[1]
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewExampleServiceClient(conn)
	for i := int64(0); i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		result, err := client.ExampleCall(ctx, &pb.ExampleRequest{Name: name, Id: i})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		log.Printf("Message: %s Id: %d", result.GetMessage(), result.GetId())
		time.Sleep(2 * time.Second)
	}
}
