package otpRepository

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
)

type OTPRepository interface {
	SaveOTP(email string) (*models.OTP, *response.ErrorResponse)
	VerifyOTP(email string, otp string) (bool, *response.ErrorResponse)
	DeleteOTP(email string) (*response.ErrorResponse)
}

