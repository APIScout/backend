package controller

import (
	"context"
	"log"
	"net/http"

	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"backend/app/internal/models"
	"backend/app/internal/mongodb"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SyncSpecificationsHandler(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		documents, err := mongodb.RetrieveDocuments(mongoClient.Database("apis"), bson.D{{}}, "specifications")

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, "Something went wrong, try again later")
			return
		}

		current := 1
		total := 1422195

		for documents.Next(context.TODO()) {
			log.Printf("Saving document %d/%d - [%d%%]", current, total, current / total)

			var document models.MongoResponse
			specification := documents.Current.Lookup("api")
			err := documents.Decode(&document)

			if err != nil {
				panic(err)
			}

			document.InitObject()

			var embeddings *models.EmbeddingResponse
			embeddings, err = embedding.PerformPipeline([]string{specification.String()}, false)

			if err != nil {
				panic(err)
			}

			var esDocument models.EsRequest
			esDocument.MongoDocument = document
			esDocument.Embedding = embeddings.Predictions[0]

			err = elastic.InsertDocument(elasticClient, esDocument, "apis")

			if err != nil {
				panic(err)
			}

			current++
		}
	}

	return fn
}
