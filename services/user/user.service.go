package userService

import (
	"auth-service-go-api/dto/response"
)

type UserService interface {
	// GetUserByEmail retrieves a user by their email address.
	GetUserByEmail(email string) (*response.UserDto, *response.ErrorResponse)

	//verifyUser 
	VerifyUser(email string) (*response.UserDto, *response.ErrorResponse)


}