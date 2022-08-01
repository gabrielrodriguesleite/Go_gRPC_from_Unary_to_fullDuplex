package main

import (
	"context"
	"fmt"
	pb "go-grpc/pb"
	"log"
	"net"

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
