package main

import (
	"backend/app/internal/controller"
	"fmt"
	"log"

	_ "backend/app/docs"
	"backend/app/internal/helpers"
	"github.com/gin-gonic/gin"
)

//	@title			API Scout
//	@version		1.0
//	@description	This is the backend for the API Scout platform.

// @BasePath	/api/v1
func main() {
	cfg := helpers.LoadConfigs()

	// Start the webserver
	router := gin.Default()
	controller.SetupRoutes(router, &cfg)
	err := router.Run(fmt.Sprintf(":%d", cfg.Backend.Port))

	if err != nil {
		log.Fatal(err)
	}
}
