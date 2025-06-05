package jwtServiceImpl

import (
	"auth-service-go-api/dto/response"
	jwtService "auth-service-go-api/services/jwt"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtServiceImpl struct {
	secretKey string
}


func NewJWTService() jwtService.JWTService {
	return &jwtServiceImpl{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

func (j *jwtServiceImpl) GenerateToken(userID uint, email string,role string) (string, *response.ErrorResponse) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
        "role":    role, // Default role, can be extended later
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiry
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", response.InternalServerError("Failed to sign token: " + err.Error())
	}

	return signedToken, nil
}

func (j *jwtServiceImpl) GenerateTemporaryToken(email string) (string, *response.ErrorResponse) {
    claims := jwt.MapClaims{
        "email": email,
        "type": "password_reset",
        "exp":  time.Now().Add(15 * time.Minute).Unix(), // 15 minutes expiration
        "iat":  time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(j.secretKey))
    if err != nil {
        return "", response.InternalServerError("Failed to generate temporary token")
    }

    return signedToken, nil
}


func (j *jwtServiceImpl) ValidateTemporaryToken(tokenString string) (string, *response.ErrorResponse) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(j.secretKey), nil
    })

    if err != nil {
        return "", response.NewUnauthorizedError("Invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", response.NewUnauthorizedError("Invalid token claims")
    }

    // Check if it's a password reset token
    tokenType, ok := claims["type"].(string)
    if !ok || tokenType != "password_reset" {
        return "", response.NewUnauthorizedError("Invalid token type")
    }

    // Get email from token
    email, ok := claims["email"].(string)
    if !ok {
        return "", response.NewUnauthorizedError("Invalid token claims")
    }

    return email, nil
}

func (j *jwtServiceImpl) ValidateToken(tokenString string) (*jwtService.TokenClaims, *response.ErrorResponse) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(j.secretKey), nil
    })

    if err != nil {
        return nil, response.NewUnauthorizedError("Invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, response.NewUnauthorizedError("Invalid token claims")
    }

    userID := uint(claims["user_id"].(float64))
    email := claims["email"].(string)

    return &jwtService.TokenClaims{
        UserID: userID,
        Email:  email,
    }, nil
}