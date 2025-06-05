package signinservice

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
)


type SigninService interface {

	SignIn(req request.SignInRequest) (*response.UserDto, *response.ErrorResponse)
}