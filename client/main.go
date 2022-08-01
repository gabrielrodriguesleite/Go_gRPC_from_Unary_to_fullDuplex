package main

import (
	"fmt"
	pb "go-grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----- MAIN -----
func main() {
	fmt.Println("Starting Client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewHeroesServiceClient(cc)
	// fmt.Println(c) // for previous test of server/client connection

	// ----- Unary -----
	// doUnary(c)

	// ----- Server Streaming -----
	// doServerStreaming(c)

	// ----- Client Streaming -----
	// doClientStreaming(c)

	// ----- Full Duplex Streaming -----
	doBidirecionalStreaming(c)
}
