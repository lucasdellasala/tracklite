package main

import (
  "context"
  "log"
  "sync"
  "time"

  pb "tracklite/api/proto"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

func main() {
  conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatalf("No se pudo conectar: %v", err)
  }
  defer conn.Close()

  client := pb.NewTrackerServiceClient(conn)

  totalRequests := 1000
  var wg sync.WaitGroup
  wg.Add(totalRequests)

  start := time.Now()

  for i := 0; i < totalRequests; i++ {
    go func(i int) {
      defer wg.Done()

      ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
      defer cancel()

      _, err := client.SendLocation(ctx, &pb.LocationUpdate{
        DeviceId:      "device-" + time.Now().String(),
        Latitude:      -34.603 + float64(i)/1000,
        Longitude:     -58.381 + float64(i)/1000,
        Timestamp:     time.Now().Unix(),
        BatteryLevel:  50.0 + float32(i%50),
        SignalQuality: "4G",
        GpsAccuracy:   5.0,
      })

      if err != nil {
        log.Printf("Error en request %d: %v", i, err)
      }
    }(i)
  }

  wg.Wait()
  duration := time.Since(start)

  log.Printf("🚀 %d peticiones enviadas en %v\n", totalRequests, duration)
}
