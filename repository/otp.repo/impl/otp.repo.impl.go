package otpRepositoryImpl

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
	otpRepository "auth-service-go-api/repository/otp.repo"
	"math/rand"

	"gorm.io/gorm"
)

type otpRepositoryImpl struct {
	DB *gorm.DB
}

// deleteOTP implements otpRepository.OTPRepository.
func (o *otpRepositoryImpl) DeleteOTP(email string) *response.ErrorResponse {
    var otpRecord models.OTP
    // Find the OTP record by email
    if err := o.DB.Where("email = ?", email).First(&otpRecord).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return response.NewBadRequestError("OTP not found for the given email")
        }
        return response.InternalServerError("Failed to delete OTP")
    }

    // Delete the OTP record
    if err := o.DB.Delete(&otpRecord).Error; err != nil {
        return response.InternalServerError("Failed to delete OTP")
    }

    return nil
}

// VerifyOTP implements otpRepository.OTPRepository.
func (o *otpRepositoryImpl) VerifyOTP(email string, otp string) (bool, *response.ErrorResponse) {
	var otpRecord models.OTP
	// Check if OTP exists for the given email
	if err := o.DB.Where("email = ? AND otp = ?", email, otp).First(&otpRecord).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, response.NewBadRequestError("Invalid OTP")
		}
		return false, response.InternalServerError("Failed to verify OTP")
	}

	// If OTP is found, return true
	return true, nil
}

// SaveOTP implements otpRepository.OTPRepository.
func (o *otpRepositoryImpl) SaveOTP(email string) (*models.OTP, *response.ErrorResponse) {
    // Generate new 6-digit OTP
    newOTPValue := uint(rand.Intn(900000) + 100000)

    // Try to find existing OTP for this email
    var existingOTP models.OTP
    result := o.DB.Where("email = ?", email).First(&existingOTP)

    if result.Error == gorm.ErrRecordNotFound {
        // Create new OTP if not found
        newOTP := &models.OTP{
            OTP:   newOTPValue,
            Email: email,
        }
        if err := o.DB.Create(newOTP).Error; err != nil {
            return nil, response.InternalServerError("Failed to create new OTP")
        }
        return newOTP, nil
    } else if result.Error != nil {
        return nil, response.InternalServerError("Database error")
    }

    // Update existing OTP
    existingOTP.OTP = newOTPValue
    if err := o.DB.Save(&existingOTP).Error; err != nil {
        return nil, response.InternalServerError("Failed to update OTP")
    }

    return &existingOTP, nil
}

func NewOTPRepositoryImpl(db *gorm.DB) otpRepository.OTPRepository {
	return &otpRepositoryImpl{db}
}
