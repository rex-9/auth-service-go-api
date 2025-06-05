package router

import (
	"auth-service-go-api/controller"
	"auth-service-go-api/helper"
	jwtService "auth-service-go-api/services/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(
	router *gin.RouterGroup,
	passwordController *controller.PasswordController,
	jwtService *jwtService.JWTService,
) {
	user := router.Group("/user")
	{
		protected := user.Group("")
        protected.Use(helper.AuthMiddleware(*jwtService))
        {
			protected.PUT("/update-password", passwordController.UpdatePassword)
		}
		// Add more user routes here
	}
}