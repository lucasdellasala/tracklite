package worker

import (
  "context"
  "encoding/json"
  "log"

  tracklite "tracklite/api/proto"
  "github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

var JobChannel = make(chan *tracklite.LocationUpdate, 100)

type LocationData struct {
  Latitude      float64 `json:"latitude"`
  Longitude     float64 `json:"longitude"`
  Timestamp     int64   `json:"timestamp"`
  BatteryLevel  float32 `json:"battery_level"`
  SignalQuality string  `json:"signal_quality"`
  GPSAccuracy   float32 `json:"gps_accuracy"`
}

func InitRedis() {
  redisClient = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
  })

  if err := redisClient.Ping(context.Background()).Err(); err != nil {
    log.Fatalf("Error connecting to Redis: %v", err)
  }
}

func worker(id int, jobs <-chan *tracklite.LocationUpdate) {
  ctx := context.Background()

  for job := range jobs {
    log.Printf("Worker %d processing job: %+v\n", id, job)

    location := LocationData{
      Latitude:      job.Latitude,
      Longitude:     job.Longitude,
      Timestamp:     job.Timestamp,
      BatteryLevel:  job.BatteryLevel,
      SignalQuality: job.SignalQuality,
      GPSAccuracy:   job.GpsAccuracy,
    }

    data, err := json.Marshal(location)
    if err != nil {
      log.Printf("Error serializing JSON: %v\n", err)
      continue
    }

    key := "device:" + job.DeviceId + ":last_position"
    err = redisClient.Set(ctx, key, data, 0).Err()
    if err != nil {
      log.Printf("Error saving to Redis: %v\n", err)
      continue
    }
  }
}

func StartWorkerPool(workerCount int) {
  InitRedis()
  for i := 1; i <= workerCount; i++ {
    go worker(i, JobChannel)
  }
}
