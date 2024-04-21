package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// TODO: Move SyncSpecificationsHandler function somewhere else
	fn := func(ctx *gin.Context) {
		documents, err := mongodb.RetrieveDocuments(mongoClient.Database("apis"), bson.D{{}}, "specifications")

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		current := 1
		total := 1422195

		for documents.Next(context.TODO()) {
			log.Printf("Saving document %d/%d - [%d%%]", current, total, current/total*100)

			var document models.MongoResponse
			specification := documents.Current.Lookup("api")
			err := documents.Decode(&document)

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			log.Printf("Mongo ID: %s", document.MongoId)

			mongoDocument := document.InitObject()
			query := fmt.Sprintf(
				`{"query": {"nested": {"path": "metadata", "query": {"match": {"metadata.mongo-id": "%s"}}}}}`,
				mongoDocument.MongoId)
			res, err := elastic.SearchDocument(elasticClient, query, "apis")

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			if len(res.Hits.Hits) == 0 {
				var embeddings *models.EmbeddingResponse
				embeddings, length, err := embedding.PerformPipeline([]string{specification.String()}, false)

				if err != nil {
					NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
					return
				}

				if len(embeddings.Predictions) != 0 {
					var esDocument models.EsRequest
					esDocument.MongoDocument = mongoDocument
					esDocument.MongoDocument.Length = length
					esDocument.Embedding = embeddings.Predictions[0]

					id, err := primitive.ObjectIDFromHex(mongoDocument.MongoId)

					if err != nil {
						NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
						return
					}

					var metricsDocument models.Metrics

					database := mongoClient.Database("apis")
					metrics, err := mongodb.RetrieveDocument(database, bson.M{"_id": id}, "metrics")

					if err != nil {
						log.Print("No metrics found for this document")
					}

					err = bson.Unmarshal(metrics, &metricsDocument)

					if err != nil {
						NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
						return
					}

					esDocument.MongoDocument.Metrics = metricsDocument
					err = elastic.InsertDocument(elasticClient, esDocument, "apis")

					if err != nil {
						NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
						return
					}
				} else {
					log.Print("No embedding was produced, skipping")
				}
			} else {
				log.Printf("Already exists, skipping")
			}

			current++
		}
	}

	return fn
}
