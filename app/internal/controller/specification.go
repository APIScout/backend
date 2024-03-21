package controller

import (
	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"github.com/goccy/go-json"
	"net/http"

	"backend/app/internal/models"
	"backend/app/internal/mongodb"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetSpecificationHandler godoc
//
//	@Summary		Get OpenAPI specification
//	@Description	Retrieve a specific OpenAPI specification's content given a valid ID
//	@Tags			specification
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Specification ID"
//	@Success		200	{object}	map[string]interface{}
//	@Failure		400	{object}	models.HTTPError
//	@Failure		500	{object}	models.HTTPError
//	@Router			/specification/{id} [get]
func GetSpecificationHandler(mongoClient *mongo.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		id := ctx.Param("id")
		db := mongoClient.Database("apis")
		objId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, "The id has not been correctly formatted")
			return
		}

		// Retrieve document based on its id
		query := bson.M{"_id": objId}
		specDoc, err := mongodb.SearchDocument(db, query, "specifications")

		if err != nil {
			NewHTTPError(ctx, http.StatusNotFound, "The document with the given ID has not been found")
			return
		}

		// Unmarshal raw bson into json document
		var jsonMap map[string]interface{}
		err = json.Unmarshal([]byte(specDoc.String()), &jsonMap)

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
			return
		}

		// Return the JSON representation of the document
		ctx.JSON(http.StatusOK, jsonMap)
	}

	return fn
}

// PostSpecificationHandler godoc
//
//	@Summary		Insert OpenAPI specifications
//	@Description	Insert new OpenAPI specifications in the database.
//	@Tags			specification
//	@Accept			json
//	@Produce		json
//	@Param			specifications	body		models.SpecificationsRequest	true	"New Specifications"
//	@Success		200				{object}	map[string]interface{}
//	@Failure		400				{object}	models.HTTPError
//	@Failure		500				{object}	models.HTTPError
//	@Router			/specification [post]
func PostSpecificationHandler(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var body models.SpecificationsRequest
		var specifications []string
		var specificationJSONs []interface{}

		err := ctx.BindJSON(&body)

		if err != nil {
			NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
			return
		}

		for index := range body.Specifications {
			var specification interface{}

			if body.Specifications[index]["api"] == nil {
				NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
				return
			}

			jsonBody, err := json.Marshal(body.Specifications[index])
			err = json.Unmarshal(jsonBody, &specification)

			if err != nil {
				NewHTTPError(ctx, http.StatusBadRequest, "The query has not been correctly formatted")
				return
			}

			specificationJSONs = append(specificationJSONs, specification)
			specifications = append(specifications, string(jsonBody))
		}

		db := mongoClient.Database("apis")
		documentIDs, err := mongodb.InsertDocuments(db, specificationJSONs, "test")

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
			return
		}

		var embeddings *models.EmbeddingResponse
		embeddings, err = embedding.PerformPipeline(specifications, false)

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
			return
		}

		for index, embeddingVal := range embeddings.Predictions {
			var request models.EsRequest
			request.MongoId = documentIDs.InsertedIDs[index].(primitive.ObjectID).Hex()
			request.Embedding = embeddingVal

			err := elastic.InsertDocument(elasticClient, request, "test")

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
				return
			}
		}
	}


	return fn
}
