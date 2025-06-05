package main

import (
	"auth-service-go-api/initializers"
	"auth-service-go-api/router"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	 services := initializers.InitializeServices()
   
	controllers, err := initializers.InitializeControllers(services)
	if err != nil {
		log.Fatal("Failed to initialize controllers:", err)
	}

	r := gin.Default()

	routerInit := router.NewRouter(&controllers)
	routerInit.SetupRoutes(r, &controllers,&services.JWTService)
	
	r.Run()
}
