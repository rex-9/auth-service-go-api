package router

import (
	"auth-service-go-api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(
    router *gin.RouterGroup,
    controller *controller.AuthController,
    otpController *controller.OTPController,
    passwordController *controller.PasswordController,
    ) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controller.Register)
        auth.POST("/verify", otpController.VerifyOTP)
        auth.POST("/signin", controller.SignIn)
        auth.POST("/send-otp", otpController.SendOTP)
        auth.POST("/forget-password-otp-verify", otpController.VerifyOTPForForgetPassword)
        auth.POST("/reset-password",passwordController.ResetPassword )
        // Add more auth routes here
    }
}