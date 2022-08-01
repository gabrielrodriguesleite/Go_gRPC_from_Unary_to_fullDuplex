package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
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
			log.Printf("Error reading client stream %v\n", err)
		}

		result += req.GetCalling().GetHero() + " - "
	}

}

func main() {
	fmt.Println("Startin server...")

	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHeroesServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
