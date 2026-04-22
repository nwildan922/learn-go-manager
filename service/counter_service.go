package service

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/nwildan922/learn-go-manager/dto"
	"github.com/nwildan922/learn-go-manager/pkg/grpc"
	redis "github.com/nwildan922/learn-go-manager/pkg/redis"
)

type CounterService struct {
	redis   *redis.RedisClient
	counter *grpc.CounterClient
}

func NewCounterService(r *redis.RedisClient, c *grpc.CounterClient) *CounterService {
	return &CounterService{
		redis:   r,
		counter: c,
	}
}

func (s *CounterService) Count() dto.BaseResponse {

	ctx := context.Background()

	counter, err := s.redis.Increment(ctx, "counter")
	log.Println("counter :")
	log.Println(counter)
	if err != nil {
		log.Println("error : ")
		log.Println(err)
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Status:  "error",
			Message: "failed to increment counter",
			Data:    nil,
		}
	}

	timestamp := time.Now().Format(time.RFC3339)

	// 🔌 call gRPC
	res, err := s.counter.SendCounter(int32(counter), timestamp)
	if err != nil {
		log.Println("error : ")
		log.Println(err)
		return dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Status:  "error",
			Message: "failed to send counter",
			Data:    nil,
		}
	}
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "success",
		Data:    res,
	}
}
