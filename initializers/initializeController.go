package initializers

import (
	"auth-service-go-api/controller"
	otpRepositoryImpl "auth-service-go-api/repository/otp.repo/impl"
	passswordRepoImpl "auth-service-go-api/repository/password.repo/impl"
	userRepoImpl "auth-service-go-api/repository/user.repo/impl"
	jwtService "auth-service-go-api/services/jwt"
	jwtServiceImpl "auth-service-go-api/services/jwt/impl"
	services "auth-service-go-api/services/otp"
	otpServiceImpl "auth-service-go-api/services/otp/impl"
	passwordservice "auth-service-go-api/services/password.service"
	passwordServiceImpl "auth-service-go-api/services/password.service/impl"
	signinservice "auth-service-go-api/services/signin.service"
	signInServiceImpl "auth-service-go-api/services/signin.service/impl"
	singUpService "auth-service-go-api/services/signup.service"
	singUpServiceImpl "auth-service-go-api/services/signup.service/impl"
	userService "auth-service-go-api/services/user"
	userServiceImpl "auth-service-go-api/services/user/impl"
)

type AppControllers struct {
    AuthController *controller.AuthController
    OtpController *controller.OTPController
    PasswordController *controller.PasswordController
}

type AppServices struct {
    JWTService jwtService.JWTService
    signUpService singUpService.SignUpService
    otpService services.OTPService
    userService userService.UserService
    signInService signinservice.SigninService
    passwordService passwordservice.PasswordService
    

}

func InitializeServices() AppServices {

      // Initialize repositories
    userRepository := userRepoImpl.NewUserRepository(DB)
    otpRepository := otpRepositoryImpl.NewOTPRepositoryImpl(DB)
    passwordRepository := passswordRepoImpl.NewPasswordRepoImpl(DB)

        // Initialize services
    signUpService := singUpServiceImpl.NewEmailSignUpService(userRepository)
    otpService := otpServiceImpl.NewOTPService(otpRepository)
    userService := userServiceImpl.NewUserServiceImpl(userRepository)
    jwtService := jwtServiceImpl.NewJWTService()
    signinService := signInServiceImpl.NewSignInServiceImpl(userRepository)
    passwordService := passwordServiceImpl.NewPasswordServiceImpl(
        passwordRepository,
        userRepository)

    return AppServices{
        JWTService: jwtService,
        signUpService: signUpService,
        otpService: otpService,
        userService: userService,
        signInService: signinService,
        passwordService: passwordService,

    }
}

func InitializeControllers(services AppServices) (AppControllers, error) {
  



    // Initialize controllers
    authController := controller.NewAuthController(
       services. signUpService,
        services.otpService,
        services.signInService,
        services.JWTService,
    )
    otpController := controller.NewOTPController(
        services.otpService,
        services.userService,
        services.JWTService,
    )
    passwordContoller := controller.NewPasswordController(
        services.passwordService,
        services.JWTService,
    )
    // Return the initializmied controllers
    return AppControllers{
        AuthController: authController,
        OtpController: otpController,
        PasswordController: passwordContoller,
    }, nil
}