package main

import (
	"backend/app/internal/elastic"
	"backend/app/internal/mongodb"
	"fmt"
	"log"

	"backend/app/internal/helpers"
	"backend/app/internal/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	cfg := helpers.LoadConfigs()

	mongoClient := mongodb.Connect(cfg.Mongo)
	elasticClient := elastic.Connect(cfg.Elastic)
	// Run the sync pipeline on a different goroutine
	go mongodb.WatchDatabase(mongoClient, elasticClient, "insert")

	router := gin.Default()
	routes.InitRoutes(router)
	err := router.Run(fmt.Sprintf(":%d", cfg.Backend.Port))

	if err != nil {
		log.Fatal(err)
	}
}
