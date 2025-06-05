package repoImpl

import (
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
	repository "auth-service-go-api/repository/user.repo"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// GetUserByEmail implements repository.UserRepo.
func (u *userRepository) GetUserByEmail(email string) (*response.UserDto, error) {
	 var user models.User
    if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    
    return &response.UserDto{
        ID:    user.ID,
        Email: user.Email,
        Password: user.Password,
        Username: user.UserName,
        IsVerified: user.ISVerified,
    }, nil
}

// ExistByEmail implements repository.UserRepo.
func (u *userRepository) ExistByEmail(email string) bool {
	var count int64
	u.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// Create implements repository.UserRepo.
func (u *userRepository) Create(user models.User) (*response.UserDto, error) {
    result := u.db.Create(&user)
    if result.Error != nil {
        return nil, result.Error
    }
       return &response.UserDto{
        ID:    user.ID,
        Email: user.Email,
        Username: user.UserName,
        Role: string(user.Role),
    }, nil
}

func (u *userRepository) UpdateVerificationStatus(email string, isVerified bool) (*response.UserDto, error) {
    var user models.User
    result := u.db.Model(&user).
        Where("email = ?", email).
        Update("is_verified", isVerified)
    
    if result.Error != nil {
        return nil, result.Error
    }

    // Get updated user
    if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }

    return &response.UserDto{
        ID:         user.ID,
        Email:      user.Email,
        IsVerified: user.ISVerified,
        Username: user.UserName,
        Role:      string(user.Role),
    }, nil
}

func NewUserRepository(db *gorm.DB) repository.UserRepo {
	return &userRepository{db}
}
