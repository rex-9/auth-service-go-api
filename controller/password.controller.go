package controller

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
	jwtService "auth-service-go-api/services/jwt"
	passwordService "auth-service-go-api/services/password.service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordController struct{
	 passwordService passwordService.PasswordService
    jwtService     jwtService.JWTService
}

func NewPasswordController(passwordService passwordService.PasswordService, jwtService jwtService.JWTService) *PasswordController {
    return &PasswordController{
        passwordService: passwordService,
        jwtService:     jwtService,
    }
}

func (c *PasswordController) ResetPassword(ctx *gin.Context) {
    var req request.ResetPasswordRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.Error(response.NewBadRequestError("Invalid request body",))
        return
    }

    // Validate reset token
	email, err := c.jwtService.ValidateTemporaryToken(req.ResetToken)
	if err != nil {
		ctx.Error(err)
		return
	}

	fmt.Println("Email from token:", email)

    if err := c.passwordService.ResetPassword(email,req.NewPassword); 
	err != nil {
        ctx.Error(err)
        return
    }

    ctx.JSON(http.StatusOK, response.ToSuccessResponseWrapper(gin.H{
        "message": "Password reset successful",
    }))
}

func (c *PasswordController) UpdatePassword(ctx *gin.Context) {
	var req request.UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(response.NewBadRequestError("Invalid request body"))
		return
	}

	userID := ctx.MustGet("userID").(uint)
    email := ctx.MustGet("email").(string)

	if _,err := c.passwordService.UpdatePassword(userID, email, req.OldPassword, req.NewPassword); 
	err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.ToSuccessResponseWrapper(gin.H{
		"message": "Password updated successfully",
	}))
}