package helper 

import (
	"auth-service-go-api/dto/response"
	jwtService "auth-service-go-api/services/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService jwtService.JWTService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.Error(response.NewUnauthorizedError("No authorization header"))
            c.Abort()
            return
        }

        // Check Bearer token format
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.Error(response.NewUnauthorizedError("Invalid token format"))
            c.Abort()
            return
        }

        // Validate the token
        claims, err := jwtService.ValidateToken(parts[1])
        if err != nil {
            c.Error(err)
            c.Abort()
            return
        }

        // Set user information in context
        c.Set("userID", claims.UserID)
        c.Set("email", claims.Email)

        c.Next()
    }
}