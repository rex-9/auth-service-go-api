package impl

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
	otpRepository "auth-service-go-api/repository/otp.repo"
	services "auth-service-go-api/services/otp"
)

type oTPServiceImpl struct {
	otpRepo otpRepository.OTPRepository
}

// DeleteOTP implements services.OTPService.
func (o *oTPServiceImpl) DeleteOTP(email string) *response.ErrorResponse {
	// Call the repository to delete the OTP
	err := o.otpRepo.DeleteOTP(email)
	if err != nil {
		return err
	}
	return nil
}

// VerifyOTP implements services.OTPService.
func (o *oTPServiceImpl) VerifyOTP(email string, otp string) (bool, *response.ErrorResponse) {
	// Call the repository to verify the OTP
	isValid, err := o.otpRepo.VerifyOTP(email, otp)
	if err != nil {
		return false, err
	}
	if !isValid {
		return false, response.NewBadRequestError("Invalid OTP")
	}

	return true, nil
}

// GenerateOTPandSave implements services.OTPService.
func (o *oTPServiceImpl) GenerateOTPandSave(email string) (*models.OTP, *response.ErrorResponse) {
	return o.otpRepo.SaveOTP(email)
}

func NewOTPService(otpRepo otpRepository.OTPRepository) services.OTPService {
	return &oTPServiceImpl{otpRepo: otpRepo}
}
