package main

import (
	"context"
	"log"
	"time"

	pb "tracklite/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
  conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTrackerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.SendLocation(ctx, &pb.LocationUpdate{
		DeviceId:      "test-device-123",
		Latitude:      -34.603722,
		Longitude:     -58.381592,
		Timestamp:     time.Now().Unix(),
		BatteryLevel:  75.5,
		SignalQuality: "4G",
		GpsAccuracy:   5.0,
	})

	if err != nil {
		log.Fatalf("Error calling RPC: %v", err)
	}

	log.Printf("Server response: %+v\n", resp)
}
