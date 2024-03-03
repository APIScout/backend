package routes

import (
	"log"
	"net/http"

	"backend/app/internal/embedding"
	"backend/app/internal/structs"

	"github.com/gin-gonic/gin"
)


type EmbeddingRequest = structs.EmbeddingRequest

func InitRetrieverRoutes(router *gin.Engine) {
	router.POST("/search", Search)
}

func Search(c *gin.Context) {
	var body EmbeddingRequest
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	embeddings := embedding.PerformPipeline([]string{body.Fragment}, true)
	log.Print(embeddings)
}
