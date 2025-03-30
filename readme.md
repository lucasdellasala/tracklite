# Tracklite 🚀

**Tracklite** is a high-efficiency, real-time location-tracking backend built in Go, designed to showcase advanced concurrency, scalability, and performance optimization techniques.

## Overview

The goal of this project is to develop a backend that efficiently handles real-time location data from mobile devices (e.g., GPS coordinates, battery level, signal quality), minimizing data and battery consumption on the client-side.

This backend supports:

- Efficient and lightweight communication using gRPC and Protocol Buffers.
- Real-time updates via WebSockets.
- Caching with Redis for fast retrieval.
- Persistent storage of historical data with configurable retention using TimescaleDB.
- Monitoring and observability using Prometheus and Grafana.

## Project Structure

```plaintext
tracklite/
├── api/
│   └── proto/
│       └── tracklite.proto
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── server/
│       └── grpc.go
├── go.mod
└── go.sum
```

## Setup and Installation

Ensure you have [Go](https://golang.org/doc/install) installed.

### Install Dependencies

```bash
go mod tidy
```

### Compile Protocol Buffers

Make sure you've installed `protoc` ([installation instructions](https://github.com/protocolbuffers/protobuf/releases)) and then run:

```bash
protoc --go_out=. --go-grpc_out=. api/proto/tracklite.proto
```

## Running the Server

To start the gRPC server, run:

```bash
go run cmd/server/main.go
```

You should see:

```
🚀 gRPC server running on port :50051
```

## Testing the Server

A simple test client is provided to simulate device messages.

Run the client:

```bash
go run client.go
```

The server logs the incoming location updates, and the client should display a confirmation:

```
Response from server: success:true message:"Location received successfully."
```

## Next Steps

- Implement Redis caching for faster access.
- Integrate WebSockets for real-time data streaming.
- Set up TimescaleDB for historical data storage.
- Add observability using Prometheus and Grafana.

## Contributing

Feel free to contribute by submitting issues, suggestions, or pull requests.

---

Built with ❤️ and Go.
