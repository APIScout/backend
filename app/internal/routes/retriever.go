package routes

import (
	"log"

	embedding "backend/app/internal/doc-embedding"

	"github.com/gin-gonic/gin"
)


func InitRetrieverRoutes(router * gin.Engine) {
	router.POST("/search", Retrieve)
}

func Retrieve(c *gin.Context) {
	fragment := c.PostForm("fragment")

	embeddings := embedding.Embed([]string{fragment})
	log.Print(embeddings)
}
