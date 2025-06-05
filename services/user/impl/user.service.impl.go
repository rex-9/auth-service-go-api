package userServiceImpl

import (
	"auth-service-go-api/dto/response"
	repository "auth-service-go-api/repository/user.repo"
	userService "auth-service-go-api/services/user"
	"net/http"

)

type userServiceImpl struct {
	UserRepository repository.UserRepo
}



// GetUserByEmail implements userService.UserService.
func (u *userServiceImpl) GetUserByEmail(email string) (*response.UserDto, *response.ErrorResponse) {
	// Retrieve user by email
	user, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, &response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Details: err.Error(),
		}
	}

	return user, nil
}

// VerifyUser implements userService.UserService.
func (u *userServiceImpl) VerifyUser(email string) (*response.UserDto, *response.ErrorResponse) {

	// Update user verification status
	updatedUser, err := u.UserRepository.UpdateVerificationStatus(email, true)
	if err != nil {
		return nil, &response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: "Failed to verify user",
			Details: err.Error(),
		}
	}

	return updatedUser, nil
}

func NewUserServiceImpl(userRepository repository.UserRepo) userService.UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}
