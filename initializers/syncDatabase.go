package initializers

import "auth-service-go-api/models"

func SyncDatabase(){

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.OTP{})

}