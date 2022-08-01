package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doUnary(c pb.HeroesServiceClient) { // add for Unary
	log.Println("Starting Unary RPC...")

	req := &pb.CallRequest{
		Calling: &pb.Calling{
			Hero: "Spiderman",
		},
	}

	res, err := c.Call(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	log.Printf("Response from Call: %v\n", res.Result)
}

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
	doUnary(c)
}
