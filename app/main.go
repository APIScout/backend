package main

import (
	"backend/app/internal/controller"
	"fmt"
	"log"

	_ "backend/app/docs"
	"backend/app/internal/elastic"
	"backend/app/internal/helpers"
	"backend/app/internal/mongodb"

	"github.com/gin-gonic/gin"
)

//	@title			API Scout
//	@version		1.0
//	@description	This is the backend for the API Scout platform.

//	@BasePath	/api/v1
func main() {
	cfg := helpers.LoadConfigs()

	mongoClient := mongodb.Connect(cfg.Mongo)
	elasticClient := elastic.Connect(cfg.Elastic)
	// Run the sync pipeline on a different goroutine
	go mongodb.WatchDatabase(mongoClient, elasticClient, "insert")
	go mongodb.WatchDatabase(mongoClient, elasticClient, "delete")


	// Start the webserver
	router := gin.Default()
	controller.SetupRoutes(router)
	err := router.Run(fmt.Sprintf(":%d", cfg.Backend.Port))

	if err != nil {
		log.Fatal(err)
	}
}
