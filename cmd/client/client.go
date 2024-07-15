package main

import (
	"context"
	"log"

	pb "github.com/V-Ader/gRPC-Semaphore/api"
	"google.golang.org/grpc"
)

func increase(client pb.SemaphoreServiceClient, value int32) {
	increaseReq := &pb.IncreaseRequest{Value: value}
	increaseResp, err := client.Increase(context.Background(), increaseReq)
	if err != nil {
		log.Fatalf("could not increase semaphore: %v", err)
	}
	log.Printf("Increase Response: %s", increaseResp.GetMessage())
}

func decrease(client pb.SemaphoreServiceClient, value int32) {
	decreaseReq := &pb.DecreaseRequest{Value: value}
	decreaseResp, err := client.Decrease(context.Background(), decreaseReq)

	if err != nil {
		log.Fatalf("could not decrease semaphore: %v", err)
	}
	log.Printf("Decrease Response: %s", decreaseResp.GetMessage())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSemaphoreServiceClient(conn)

	increase(client, 10)

	decrease(client, 4)

}
