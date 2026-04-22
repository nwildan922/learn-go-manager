package service

import (
	"net/http"

	"github.com/nwildan922/learn-go-manager/dto"
)

func HealthCheck(message string) dto.BaseResponse {
	return dto.BaseResponse{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: message,
		Data:    nil,
	}
}
