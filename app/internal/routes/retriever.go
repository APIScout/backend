package routes

import (
	"log"
	"net/http"

	embedding "backend/app/internal/doc-embedding"
	"backend/app/internal/structs"

	"github.com/gin-gonic/gin"
)


type EmbeddingRequest = structs.EmbeddingRequest

func InitRetrieverRoutes(router *gin.Engine) {
	router.POST("/search", Retrieve)
}

func Retrieve(c *gin.Context) {
	var body EmbeddingRequest
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	embeddings := embedding.EmbedFragment([]string{body.Fragment}, true)
	log.Print(embeddings)
}
