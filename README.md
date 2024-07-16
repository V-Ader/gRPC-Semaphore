# gRPC Semaphore Service

This project implements a gRPC-based semaphore service in Go. The semaphore service allows clients to increase and decrease a counting semaphore. The service is implemented using gRPC and Protobuf for RPC communication.

## Project Structure

```grpc-semaphore/
├── api/
│ ├── semaphore.proto
│ ├── semaphore.pb.go
│ └── semaphore_grpc.pb.go
├── cmd/
│ ├── server/
│ │ └── server.go
│ └── client/
│   └── client.go
├── pkg/
│ └── semaphore/
│ ├── semaphore.go
│ ├── semaphore_server.go
│ └── semaphore_server_test.go
├── go.mod
└── go.sum
```


## Requirements

- Go 1.21 or later
- Protocol Buffers compiler (`protoc`)
- gRPC Go library

## Setup

1. **Install Go**:
    - Follow the installation instructions at [golang.org](https://golang.org/doc/install).

2. **Install Protocol Buffers Compiler**:
    - Download and install `protoc` from the [Protocol Buffers GitHub repository](https://github.com/protocolbuffers/protobuf/releases).

3. **Install gRPC and Protobuf Go Plugins**:
    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

4. **Clone the repository**:
    ```sh
    git clone https://github.com/V-Ader/gRPC-Semaphore.git
    cd grpc-semaphore
    ```

5. **Generate gRPC code from Protobuf definitions**:
    ```sh
    cd api
    protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative semaphore.proto
    ```

## Running the Server

To start the gRPC server:

```sh
cd cmd/server
go run main.go
```
```
The server will start and listen on port 50051.
```

Running the Client
To run the client:
```sh
cd cmd/client
go run main.go
```

The client will perform increase and decrease operations on the semaphore and print the responses.


## Testing

Unit tests for the semaphore logic are located in the pkg/semaphore directory. To run the unit tests:
```sh
go test ./pkg/semaphore/...
```
## Usage
Increase Semaphore
To increase the semaphore:
```go
err := IncreaseSemaphore(client, 2)
if err != nil {
    log.Fatalf("could not increase semaphore: %v", err)
}
```
Decrease Semaphore
To decrease the semaphore:

```go
err := DecreaseSemaphore(client, 1)
if err != nil {
    log.Fatalf("could not decrease semaphore: %v", err)
}
```


