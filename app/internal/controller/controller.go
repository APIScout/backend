package controller

import (
	"backend/app/internal/elastic"
	"backend/app/internal/models"
	"backend/app/internal/mongodb"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, config *models.Config) {
	mongoClient := mongodb.Connect(config.Mongo)
	elasticClient := elastic.Connect(config.Elastic)

	// Create routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/search", SearchHandler(mongoClient, elasticClient))
		v1.POST("/preprocess", GetEmbedding)
		spec := v1.Group("/specification")
		{
			spec.POST("/", PostSpecificationHandler(mongoClient, elasticClient))
			spec.GET("/:id", GetSpecificationHandler(mongoClient, elasticClient))
			spec.PUT("/sync", SyncSpecificationsHandler(mongoClient, elasticClient))
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func NewHTTPError(ctx *gin.Context, status int, message string) {
	body := models.HTTPError{
		Code:    status,
		Message: message,
	}

	ctx.AbortWithStatusJSON(status, body)
}
