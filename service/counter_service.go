package service

import (
	"context"
	"log"
	"net/http"

	"github.com/nwildan922/learn-go-manager/dto"
	redis "github.com/nwildan922/learn-go-manager/pkg/redis"
)

type CounterService struct {
	redis *redis.RedisClient
}

func NewCounterService(r *redis.RedisClient) *CounterService {
	return &CounterService{
		redis: r,
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

	return dto.BaseResponse{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "success",
		Data:    counter,
	}
}
