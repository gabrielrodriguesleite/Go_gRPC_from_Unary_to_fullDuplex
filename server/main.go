package main

import (
	"fmt"
	pb "go-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ----- Main -----
func main() {

	l, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHeroesServiceServer(s, &server{})

	fmt.Println("Startin server...")

	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
