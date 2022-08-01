package main

import (
	"context"
	pb "go-grpc/pb"
	"io"
	"log"
	"time"
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
// func doServerStreaming(c pb.HeroesServiceClient) {
// 	log.Printf("Starting Server Straming RPC...")

// 	req := &pb.CallTeamRequest{
// 		Calling: &pb.Calling{
// 			Hero: "X-Men",
// 		},
// 	}

// 	stream, err := c.CallTeam(context.Background(), req)

// 	if err != nil {
// 		log.Fatalf("Error calling Server Streaming CallTeam RPC: %v", err)
// 	}

// 	for { // while true ?
// 		res, err := stream.Recv()

// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			log.Printf("Error while reading stream: %v", err)
// 		}

// 		log.Printf("Response from CallTeam: %v\n", res.GetResult())
// 	}

// }

// ----- Client Streaming -----
// func doClientStreaming(c pb.HeroesServiceClient) {
// 	log.Println("Starting Client Streaming RPC...")

// 	stream, err := c.CallManyHeroes(context.Background())
// 	if err != nil {
// 		log.Printf("Error while calling CallManyHeroes: %v", err)
// 	}

// 	requests := []*pb.CallManyHeroesRequest{
// 		{
// 			Calling: &pb.Calling{
// 				Hero: "Mister Fantastic",
// 			},
// 		}, {
// 			Calling: &pb.Calling{
// 				Hero: "Invisible Woman",
// 			},
// 		}, {
// 			Calling: &pb.Calling{
// 				Hero: "Thing",
// 			},
// 		}, {
// 			Calling: &pb.Calling{
// 				Hero: "Human Torch",
// 			},
// 		},
// 	}

// 	for _, req := range requests {
// 		stream.Send(req)
// 		time.Sleep(1 * time.Second) // delay for suspense
// 	}

// 	res, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Printf("Error while receiving response from CallManyHeroes: %v", err)
// 	}

// 	log.Printf("CallManyHeroes response: %v", res)
// }

// Full Duplex streaming
func doBidirecionalStreaming(c pb.HeroesServiceClient) {
	log.Printf("Starting Bi Direcional Streaming RPC...")

	stream, err := c.CallEveryone(context.Background())
	if err != nil {
		log.Printf("Error while calling CallEveryoneHeroes: %v", err)
	}

	requests := []*pb.CallEveryoneRequest{
		{
			Calling: &pb.Calling{
				Hero: "Conan",
			},
		}, {
			Calling: &pb.Calling{
				Hero: "Bêlit",
			},
		}, {
			Calling: &pb.Calling{
				Hero: "Red Sonja",
			},
		},
	}

	// Concurrency
	waitChannel := make(chan struct{})

	// send the messages using Go Routine
	go func() {
		for _, req := range requests {
			log.Printf("Sending req: %v", req)
			stream.Send(req)
			time.Sleep(2 * time.Second) // delay for suspense
		}
		stream.CloseSend()
	}()

	// receive messages using Go Routine
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving response from CallEveryone: %v", err)
				break
			}

			result := res.GetResult()
			log.Printf("Call Everyone response: %v\n", result)
		}
		close(waitChannel)
	}()

	// block until everything is done
	<-waitChannel
}
