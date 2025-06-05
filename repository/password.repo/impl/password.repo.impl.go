package passswrdRepoImpl

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
	passwordrepo "auth-service-go-api/repository/password.repo"

	"gorm.io/gorm"
)

type PasswordRepoImpl struct {
	DB *gorm.DB
}

// updatePassword implements passwordrepo.PasswordRepository.
func (p *PasswordRepoImpl) UpdatePassword(userId uint, newPassword string) *response.ErrorResponse {
	var user models.User
	if err := p.DB.First(&user, userId).Error; err != nil {
		return &response.ErrorResponse{
			Message: "User not found",
			Code:    404,
		}
	}

	user.Password = newPassword
	if err := p.DB.Save(&user).Error; err != nil {
		return &response.ErrorResponse{
			Message: "Failed to update password",
			Code:    500,
		}
	}

	return nil
}

// resetPassword implements passwordrepo.PasswordRepository.
func (p *PasswordRepoImpl) ResetPassword(email string, newPassword string) *response.ErrorResponse {
	var user models.User
	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return &response.ErrorResponse{
			Message: "User not found",
			Code:    404,
		}
	}

	user.Password = newPassword
	if err := p.DB.Save(&user).Error; err != nil {
		return &response.ErrorResponse{
			Message: "Failed to reset password",
			Code:    500,
		}
	}

	return nil
}

func NewPasswordRepoImpl(db *gorm.DB) passwordrepo.PasswordRepository {
	return &PasswordRepoImpl{
		DB: db,
	}
}
