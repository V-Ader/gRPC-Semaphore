package semaphore

import (
	"context"
	pb "github.com/V-Ader/gRPC-Semaphore/api"
)

type Server struct {
	pb.UnimplementedSemaphoreServiceServer
	Sem *Semaphore
}

func (s *Server) Increase(ctx context.Context, req *pb.IncreaseRequest) (*pb.SemaphoreResponse, error) {
	s.Sem.Increase(int(req.GetValue()))
	return &pb.SemaphoreResponse{Success: true, Message: "Semaphore increased"}, nil
}

func (s *Server) Decrease(ctx context.Context, req *pb.DecreaseRequest) (*pb.SemaphoreResponse, error) {
	err := s.Sem.Decrease(int(req.GetValue()))
	if err != nil {
		return &pb.SemaphoreResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.SemaphoreResponse{Success: true, Message: "Semaphore decreased"}, nil
}
