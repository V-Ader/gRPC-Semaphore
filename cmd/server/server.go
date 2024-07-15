package main

import (
	"log"
	"net"

	pb "github.com/V-Ader/gRPC-Semaphore/api"
	"github.com/V-Ader/gRPC-Semaphore/pkg/semaphore"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSemaphoreServiceServer(grpcServer, &semaphore.Server{
		Sem: semaphore.NewSemaphore(1),
	})

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
