package passwordservice

import "auth-service-go-api/dto/response"

type PasswordService interface {
	UpdatePassword(userId uint,email string,oldPassword string, newPassword string,) (*response.UserDto, *response.ErrorResponse)
	ResetPassword(email string, newPassword string) ( *response.ErrorResponse)
}