package serviceImpl

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
	"auth-service-go-api/models"
	repository "auth-service-go-api/repository/user.repo"
	services "auth-service-go-api/services/signup.service"

	"golang.org/x/crypto/bcrypt"
)

type emailSignUpServiceImpl struct {
    userRepo repository.UserRepo
}

func NewEmailSignUpService(userRepo repository.UserRepo) services.SignUpService {
    return &emailSignUpServiceImpl{userRepo: userRepo}
}

func (s *emailSignUpServiceImpl) SignUp(req request.SignUpRequest) (*response.UserDto, *response.ErrorResponse) {
    
	if err := s.checkEmailIsAlreadyRegistered(req); err != nil {
        return nil, err
    }

    hashedPassword, err := s.hashPassword(req.Password)
    if err != nil {
        return nil, response.InternalServerError("Failed to process password")
    }

    user := models.User{
        Email:    req.Email,
        Password: hashedPassword,
        UserName: req.UserName,
    }

   createdUser, err := s.userRepo.Create(user)
    if err != nil {
        return nil, response.InternalServerError("Failed to create user")
    }

    return &response.UserDto{
        ID:    createdUser.ID,
        Email: createdUser.Email,
        Username: createdUser.Username,
    }, nil
}

func (s *emailSignUpServiceImpl) checkEmailIsAlreadyRegistered(req request.SignUpRequest) *response.ErrorResponse {
    if s.userRepo.ExistByEmail(req.Email) {
        return response.UserAlreadyExitError()
    }
    return nil
}

func (s *emailSignUpServiceImpl) hashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}