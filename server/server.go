package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb"
	"io"
	"log"
	"time"
)

type server struct {
	pb.UnimplementedHeroesServiceServer
}

// ----- Unary -----
func (*server) Call(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	log.Printf("Greet was invoked with %v\n", req)

	hero := req.GetCalling().GetHero()

	result := hero + " is on its way!"

	res := &pb.CallResponse{
		Result: result,
	}

	return res, nil
}

// ----- Server Stream -----
func (*server) CallTeam(req *pb.CallTeamRequest, stream pb.HeroesService_CallTeamServer) error {
	log.Printf("CallTeam was invoked with: %v\n", req)

	hero := req.GetCalling().GetHero()

	xmen := []string{"Cyclops", "Iceman", "Angel", "Beast", "Phoenix"}

	for _, x := range xmen {
		result := fmt.Sprintf("You called the %v! %v is goin!!\n", hero, x)
		res := &pb.CallTeamResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second) // 1 sec for suspense
	}
	return nil
}

// ----- Client Stream -----
func (*server) CallManyHeroes(stream pb.HeroesService_CallManyHeroesServer) error {
	log.Printf("CallManyHeroes was invoked with a streaming request")

	result := "You called "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.CallManyHeroesResponse{
				Result: result + "The Fantastic Four!",
			})
		}

		if err != nil {
			log.Printf("Error reading client stream: %v\n", err)
		}

		result += req.GetCalling().GetHero() + " - "
	}

}

// ----- Full Duplex Stream -----
func (*server) CallEveryone(stream pb.HeroesService_CallEveryoneServer) error {
	log.Println("CallEveryone was invoked with a streaming request.")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("Request end")
			return nil
		}

		log.Printf("Request: %v\n", req)

		if err != nil {
			log.Printf("Error reading client stream.")
			return err
		}

		hero := req.GetCalling().GetHero()

		result := hero + " is on its way!"

		err = stream.Send(&pb.CallEveryoneResponse{
			Result: result,
		})

		if err != nil {
			log.Println("Error while sending data to client.")
			return err
		}
	}
}
