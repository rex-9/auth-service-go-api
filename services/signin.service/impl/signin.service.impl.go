package signInServiceImpl

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
	repository "auth-service-go-api/repository/user.repo"
	signinservice "auth-service-go-api/services/signin.service"

	"golang.org/x/crypto/bcrypt"
)

type signInServiceImpl struct {
	userRepository repository.UserRepo
}

// SignIn implements signinservice.SigninService.
func (s *signInServiceImpl) SignIn(req request.SignInRequest) (*response.UserDto, *response.ErrorResponse) {
	 user, err := s.userRepository.GetUserByEmail(req.Email)
    if err != nil {
        return nil, response.NewUnauthorizedError("Invalid credentials")
    }

    // Verify password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return nil, response.NewUnauthorizedError("Invalid credentials")
    }

	user.Password= "" // Clear password before returning
    return user, nil
}

func NewSignInServiceImpl(userRepository repository.UserRepo) signinservice.SigninService {
	return &signInServiceImpl{
		userRepository: userRepository,
	}
}
