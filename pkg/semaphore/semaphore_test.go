package semaphore

import (
	"context"
	"testing"

	pb "github.com/V-Ader/gRPC-Semaphore/api"
	"github.com/stretchr/testify/assert"
)

func TestServer_Increase(t *testing.T) {
	s := &Server{
		Sem: NewSemaphore(1),
	}

	req := &pb.IncreaseRequest{Value: 2}
	res, err := s.Increase(context.Background(), req)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Semaphore increased", res.Message)
	assert.Equal(t, 3, s.Sem.permits)
}

func TestServer_Decrease(t *testing.T) {
	s := &Server{
		Sem: NewSemaphore(3),
	}

	req := &pb.DecreaseRequest{Value: 2}
	res, err := s.Decrease(context.Background(), req)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "Semaphore decreased", res.Message)
	assert.Equal(t, 1, s.Sem.permits)

	// Test decrease beyond available permits
	req = &pb.DecreaseRequest{Value: 2}
	res, err = s.Decrease(context.Background(), req)
	assert.NoError(t, err)
	assert.False(t, res.Success)
	assert.Equal(t, "not enough permits: requested 2, available 1", res.Message)
	assert.Equal(t, 1, s.Sem.permits)
}
