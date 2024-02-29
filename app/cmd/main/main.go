package main

import (
	"fmt"
	"log"

	"backend/app/internal/helpers"
	"backend/app/internal/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg := helpers.LoadConfigs()

	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run(fmt.Sprintf(":%d", cfg.Backend.Port))

	if err != nil {
		log.Fatal(err)
	}
}
