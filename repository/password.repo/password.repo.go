package passwordrepo

import "auth-service-go-api/dto/response"

type PasswordRepository interface {
	ResetPassword(email string, newPassword string) (*response.ErrorResponse)
	UpdatePassword(userId uint, newPassword string) *response.ErrorResponse
}