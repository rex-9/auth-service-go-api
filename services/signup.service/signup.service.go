package services

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
)


type SignUpService interface {

    SignUp(req request.SignUpRequest) (*response.UserDto, *response.ErrorResponse)

}




