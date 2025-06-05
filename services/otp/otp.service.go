package services

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
)

type OTPService interface {
	GenerateOTPandSave(email string) (*models.OTP, *response.ErrorResponse)
	VerifyOTP(email string, otp string) (bool, *response.ErrorResponse)
	DeleteOTP(email string) (*response.ErrorResponse)
}

