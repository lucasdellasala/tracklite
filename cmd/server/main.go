package main

import (
	"log"
	"net"

	pb "tracklite/api/proto"
	"tracklite/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTrackerServiceServer(s, &server.Server{})

	log.Println("🚀 Servidor gRPC running at port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
