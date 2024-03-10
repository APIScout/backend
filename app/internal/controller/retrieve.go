package controller

import (
	"log"
	"net/http"

	"backend/app/internal/embedding"
	"backend/app/internal/models"

	"github.com/gin-gonic/gin"
)

type EmbeddingRequest = models.EmbeddingRequest

// Search godoc
//	@Summary		Search OpenAPI specifications
//	@Description	retrieve OpenAPI specifications matching the given query
//	@Tags			search
//	@Accept			json
//	@Produce		json
//	@Success		200	{string} ok
//	@Failure		400	{string} Bad Request
//	@Failure		500	{string} Internal Server Error
//	@Router			/search [post]
func Search(ctx *gin.Context) {
	var body EmbeddingRequest
	err := ctx.BindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
	}

	embeddings := embedding.PerformPipeline([]string{body.Fragment}, true)
	log.Print(len(embeddings.Predictions[0]))
}
