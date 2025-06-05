package controller

import (
	"auth-service-go-api/dto/request"
	"auth-service-go-api/dto/response"
	jwtService "auth-service-go-api/services/jwt"
	otpServices "auth-service-go-api/services/otp"
	signinservice "auth-service-go-api/services/signin.service"
	singupServices "auth-service-go-api/services/signup.service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	emailSignUpService singupServices.SignUpService
	otpService         otpServices.OTPService
	signinservice signinservice.SigninService
	jwtService jwtService.JWTService
}

// NewAuthController constructs the controller with its service.
func NewAuthController(
	emailSignUpService singupServices.SignUpService,
	otpService otpServices.OTPService,
	signInService signinservice.SigninService,
	jwtService jwtService.JWTService,
	) *AuthController {
	return &AuthController{
		emailSignUpService: emailSignUpService,
		 otpService: otpService,
		 signinservice: signInService,
		 jwtService: jwtService,
		}
}

func (ctr *AuthController) Register(c *gin.Context) {

	var req request.SignUpRequest

	//check the request is valid or not
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(response.NewBadRequestError("Invalid request body"))
		return
	}

	//sign up to the database using service
	userDto, err := ctr.emailSignUpService.SignUp(req)

	if err != nil {
		c.Error(err)
		return
	}

	//send otp for verification
	otpModel , err := ctr.otpService.GenerateOTPandSave(userDto.Email);
	if err != nil {
		c.Error(err)
		return
	}

	//actually send email service come here
	res := response.ToSuccessResponseWrapper(otpModel)

	c.JSON(http.StatusCreated, res)
}

func (ctr *AuthController) SignIn(c *gin.Context) {
	var req request.SignInRequest
	//check the request is valid or not
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(response.NewBadRequestError("Invalid request body"))
		return
	}

	//sign in to the database using service
	userDto, err := ctr.signinservice.SignIn(req)
    if err != nil {
        c.Error(err)
        return
    }
 	token, err := ctr.jwtService.GenerateToken(userDto.ID, userDto.Email,userDto.Role)
    if err != nil {
        c.Error(err)
        return
    }

	userDto.Token = token
	res := response.ToSuccessResponseWrapper(userDto)
	// fmt.Print("UserDto: ", token)

	c.JSON(200, res)
}

 
