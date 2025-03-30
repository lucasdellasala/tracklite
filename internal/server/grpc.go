package server

import (
	"context"
	"log"

	pb "tracklite/api/proto"
	"tracklite/internal/worker"
)

type Server struct {
	pb.UnimplementedTrackerServiceServer
}

func (s *Server) SendLocation(ctx context.Context, in *pb.LocationUpdate) (*pb.LocationUpdateResponse, error) {
	log.Printf("Message received: %+v\n", in)

	worker.JobChannel <- in

	return &pb.LocationUpdateResponse{
		Success: true,
		Message: "Location received correctly.",
	}, nil
}
