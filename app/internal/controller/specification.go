package controller

import (
	"backend/app/internal/mongodb"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetSpecificationHandler godoc
//	@Summary		Get OpenAPI specification
//	@Description	Retrieve a specific OpenAPI specification's content given a valid ID
//	@Tags			specification
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Specification ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	models.HTTPError
//	@Failure		500	{object}	models.HTTPError
//	@Router			/specifications/{id} [get]
func GetSpecificationHandler(mongoClient *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		id := ctx.Param("id")
		db := mongoClient.Database("apis")
		objId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, "The id has not been correctly formatted")
		}

		// Retrieve document based on its id
		query := bson.M{"_id": objId}
		specDoc, err := mongodb.SearchDocument(db, query, "specifications")

		if err != nil {
			NewHTTPError(ctx, http.StatusNotFound, "The document with the given ID has not been found")
		}

		// Unmarshal raw bson into json document
		var jsonMap map[string]interface{}
		err = json.Unmarshal([]byte(specDoc.String()), &jsonMap)

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
		}

		// Return the JSON representation of the document
		ctx.JSON(http.StatusOK, jsonMap)
	}

	return fn
}

// PostSpecificationHandler godoc
//	@Summary		Insert OpenAPI specifications
//	@Description	Insert new OpenAPI specifications in the database.
//	@Tags			specification
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Specification ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	models.HTTPError
//	@Failure		500	{object}	models.HTTPError
//	@Router			/specifications/insert [post]
func PostSpecificationHandler(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {

	}

	return fn
}
