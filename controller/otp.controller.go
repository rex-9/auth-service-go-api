package controller

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
	"auth-service-go-api/services/jwt"
	otpService "auth-service-go-api/services/otp"
	userService "auth-service-go-api/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OTPController struct {
	otpService  otpService.OTPService
	userService userService.UserService
	jwtService  jwtService.JWTService
}

func NewOTPController(otpService otpService.OTPService, userService userService.UserService, jwtService jwtService.JWTService) *OTPController {
	return &OTPController{
		otpService:  otpService,
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *OTPController) SendOTP(ctx *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(response.NewBadRequestError("Invalid email format"))
		return
	}

	otpModel, err := c.otpService.GenerateOTPandSave(body.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.ToSuccessResponseWrapper(otpModel))
}

func (c *OTPController) VerifyOTP(ctx *gin.Context) {
	if err := c.verifyOTPCommon(ctx); err != nil {
		return
	}

	req := ctx.MustGet("verifyRequest").(*request.VerifyOTPRequest)
	userDto, err := c.userService.VerifyUser(req.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	token, err := c.jwtService.GenerateToken(userDto.ID, userDto.Email,userDto.Role)
	if err != nil {
		ctx.Error(err)
		return
	}

	userDto.Token = token

	ctx.JSON(http.StatusOK, response.ToSuccessResponseWrapper(userDto))
}

func (c *OTPController) VerifyOTPForForgetPassword(ctx *gin.Context) {
	if err := c.verifyOTPCommon(ctx); err != nil {
		return
	}

	req := ctx.MustGet("verifyRequest").(*request.VerifyOTPRequest)
	tempToken, err := c.jwtService.GenerateTemporaryToken(req.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.ToSuccessResponseWrapper(gin.H{
		"message":   "OTP verified successfully",
		"tempToken": tempToken,
		"email":     req.Email,
	}))
}

// Common OTP verification logic
func (c *OTPController) verifyOTPCommon(ctx *gin.Context) error {
	var req request.VerifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(response.NewBadRequestError("Invalid request body"))
		return err
	}

	isVerified, err := c.otpService.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		ctx.Error(err)
		return err
	}

	if !isVerified {
		err := &response.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: "Invalid OTP",
			Details: "The OTP provided is invalid or has expired"}
		ctx.Error(err)
		return err
	}

	if err := c.otpService.DeleteOTP(req.Email); err != nil {
		ctx.Error(err)
		return err
	}

	ctx.Set("verifyRequest", &req)
	return nil
}
