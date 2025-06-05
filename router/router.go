package router

import (
	"auth-service-go-api/helper"
	"auth-service-go-api/initializers"
	jwtService "auth-service-go-api/services/jwt"

	"github.com/gin-gonic/gin"
)

type Router struct {
    controllers *initializers.AppControllers
}

func NewRouter(controllers *initializers.AppControllers) *Router {
    return &Router{
        controllers: controllers,
    }
}

func (r *Router) SetupRoutes(app *gin.Engine, ctr *initializers.AppControllers, jwtService *jwtService.JWTService) {

	app.Use(helper.ErrorHandler())

    api := app.Group("/api")
    {
        RegisterAuthRoutes(
            api,
            ctr.AuthController,
            ctr.OtpController,
            ctr.PasswordController,
        )

        RegisterUserRoutes(
            api,
            ctr.PasswordController,
            jwtService,
        )
    }
}