package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ----- Unary -----
// func doUnary(c pb.HeroesServiceClient) {
// 	log.Println("Starting Unary RPC...")

// 	req := &pb.CallRequest{
// 		Calling: &pb.Calling{
// 			Hero: "Spiderman",
// 		},
// 	}

// 	res, err := c.Call(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error Call RPC: %v", err)
// 	}

// 	log.Printf("Response from Call: %v\n", res.Result)
// }

// ----- Server Streaming -----
func doServerStreaming(c pb.HeroesServiceClient) {
	log.Printf("Starting Server Straming RPC...")

	req := &pb.CallTeamRequest{
		Calling: &pb.Calling{
			Hero: "X-Men",
		},
	}

	stream, err := c.CallTeam(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling Server Streaming CallTeam RPC: %v", err)
	}

	for { // while true ?
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error while reading stream: %v", err)
		}

		log.Printf("Response from CallTeam: %v\n", res.GetResult())
	}

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
	// doUnary(c)

	// ----- Server Streaming -----
	doServerStreaming(c)
}
