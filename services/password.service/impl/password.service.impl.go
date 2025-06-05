package passwordServiceImpl

import (
	"auth-service-go-api/dto/response"
	passwordrepo "auth-service-go-api/repository/password.repo"
	repository "auth-service-go-api/repository/user.repo"
	passwordservice "auth-service-go-api/services/password.service"

	"golang.org/x/crypto/bcrypt"
)

type passwordServiceImpl struct {
	passwordrepo passwordrepo.PasswordRepository
	userRepo repository.UserRepo
}

// UpdatePassword implements passwordservice.PasswordService.
func (p *passwordServiceImpl) UpdatePassword(userId uint, email string, oldPassword string, newPassword string) (*response.UserDto, *response.ErrorResponse) {
	

	user, err := p.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, response.InternalServerError("Failed to retrieve user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
        return nil, response.NewUnauthorizedError("Old password is incorrect")
    }

	hashedNewPassword, err := p.hashPassword(newPassword)
	if err != nil {
		return nil, response.InternalServerError("Failed to process new password")
	}

	updateError := p.passwordrepo.UpdatePassword(userId, hashedNewPassword)
	if updateError != nil {
		return nil,updateError
	}

	user.Password = ""
	return user, nil
}

// ResetPassword implements passwordservice.PasswordService.
func (p *passwordServiceImpl) ResetPassword(email string, password string) *response.ErrorResponse {
	hashedPassword, err := p.hashPassword(password)
	if err != nil {
		return response.InternalServerError("Failed to process password")
	}

	resetErr := p.passwordrepo.ResetPassword(email, hashedPassword)
	if resetErr != nil {
		return resetErr
	}
	return nil
}

func NewPasswordServiceImpl(passwordrepo passwordrepo.PasswordRepository,userRepository repository.UserRepo) passwordservice.PasswordService {
	return &passwordServiceImpl{
		passwordrepo: passwordrepo,
		userRepo: userRepository,
	}
}

func (p *passwordServiceImpl) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
