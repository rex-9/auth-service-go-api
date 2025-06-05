package repository

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
)

type UserRepo interface {
	Create(user models.User) (*response.UserDto, error)
	ExistByEmail(email string) bool
	GetUserByEmail(email string) (*response.UserDto, error)
	UpdateVerificationStatus(email string, status bool) (*response.UserDto, error)
}