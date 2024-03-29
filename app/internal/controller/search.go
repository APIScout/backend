package controller

import (
	"log"
	"net/http"

	"backend/app/internal/embedding"
	"backend/app/internal/models"

	"github.com/gin-gonic/gin"
)

// Search godoc
//
//		@Summary		Search OpenAPI specifications
//		@Description	Retrieve OpenAPI specifications matching the given query
//		@Tags			search
//		@Accept			json
//		@Produce		json
//		@Param			fragment	body		models.EmbeddingRequest	true	"Search query"
//	 TODO: Change this
//		@Success		200			{string}	OK
//		@Failure		400			{object}	models.HTTPError
//		@Router			/search [post]
func Search(ctx *gin.Context) {
	var body models.EmbeddingRequest
	err := ctx.BindJSON(&body)

	if err != nil {
		NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
		return
	}

	embeddings, err := embedding.PerformPipeline([]string{body.Fragment}, true)

	if err != nil {
		NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
		return
	}

	log.Print(body)
	log.Print(len(embeddings.Predictions[0]))
}
