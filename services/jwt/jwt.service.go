package jwtService

import "auth-service-go-api/dto/response"

 type TokenClaims struct {
    UserID uint   `json:"user_id"`
    Email  string `json:"email"`
}

type JWTService interface {
	GenerateToken(userID uint, email string,role string) (string, *response.ErrorResponse)
	GenerateTemporaryToken(email string) (string, *response.ErrorResponse)
	ValidateTemporaryToken(token string) (string, *response.ErrorResponse)
	 ValidateToken(tokenString string) (*TokenClaims, *response.ErrorResponse)
}

