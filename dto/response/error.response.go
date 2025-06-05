package response

import "net/http"

type ErrorResponse struct {
	Code    int
	Message string
	Details string
}

// Error implements error.
func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewBadRequestError(details string) *ErrorResponse {
	return &ErrorResponse{
		Code:    400,
		Message: "BAD_REQUEST",
		Details: details,
	}
}

func UserAlreadyExitError() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusForbidden,
		Message: "User Already Existed",
		Details: "User already registered by this Phone Number!",
	}
}

func InternalServerError(detail string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server occur",
		Details: detail,
	}
}

func NewUnauthorizedError(detail string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
		Details: detail,
	}
}
