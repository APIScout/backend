package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"backend/app/internal/models"
	"backend/app/internal/mongodb"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SyncSpecificationsHandler(mongoClient *mongo.Client, elasticClient *elasticsearch.Client) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		var documents *mongo.Cursor
		var err error

		skip := 0
		skipQuery := ctx.Query("skip")

		if strings.Compare(skipQuery, "") != 0 && strings.Compare(skipQuery, "auto") != 0 {
			skip, err = strconv.Atoi(skipQuery)

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, "the skip parameter must be a number")
				return
			}
		} else if strings.Compare(skipQuery, "auto") == 0 {
			count, err := elastic.GetDocumentCount(elasticClient, "apis")

			if err != nil {
				NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
				return
			}

			skip = int(count)
		}

		documents, err = mongodb.RetrieveDocuments(mongoClient.Database("apis"), bson.D{{}}, "specifications")

		if err != nil {
			NewHTTPError(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		current := 1
		total, err := mongoClient.Database("apis").Collection("specifications").CountDocuments(
			context.TODO(),
			bson.D{},
			options.Count().SetHint("_id_"),
		)

		if err != nil {
			panic(err)
		}

		var prec int64

		for documents.Next(context.TODO()) {
			start := time.Now().UnixNano()

			percentage := (float64(current) / float64(total)) * 100.0
			totTime := (float64(prec) * float64(total - int64(current))) / 3600000000000.0
			log.Printf("Saving document %d/%d - [%.1f%%] - Done in: %.1fh", current, total, percentage, totTime)

			if skip > current {
				log.Printf("Skip query set to %d, skipping...", skip)

				prec = time.Now().UnixNano()-start
				current++

				continue
			}

			var document models.MongoResponse
			specification := documents.Current.Lookup("api")
			err := documents.Decode(&document)

			if err != nil {
				log.Print("Error decoding document, skipping...",)
				continue
			}

			log.Printf("Mongo ID: %s", document.MongoId)

			mongoDocument := document.InitObject()
			query := fmt.Sprintf(
				`{"query": {"nested": {"path": "metadata", "query": {"match": {"metadata.mongo-id": "%s"}}}}}`,
				mongoDocument.MongoId)
			res, err := elastic.SearchDocument(elasticClient, query, "apis")

			if err != nil {
				log.Print("Error retrieving document, skipping...",)
				continue
			}

			if len(res.Hits.Hits) == 0 {
				var embeddings *models.EmbeddingResponse
				embeddings, length, err := embedding.PerformPipeline([]string{specification.String()}, false)

				if err != nil {
					log.Print("Error embedding document, skipping...",)
					continue
				}

				if len(embeddings.Predictions) != 0 {
					var esDocument models.EsRequest
					esDocument.MongoDocument = *mongoDocument
					esDocument.MongoDocument.Length = length[0]
					esDocument.Embedding = embeddings.Predictions[0]

					id, err := primitive.ObjectIDFromHex(mongoDocument.MongoId)

					if err != nil {
						log.Print("Error decoding document, skipping...",)
						continue
					}

					var metricsDocument models.Metrics

					database := mongoClient.Database("apis")
					metrics, err := mongodb.RetrieveDocument(database, bson.M{"_id": id}, "metrics")

					if err != nil {
						log.Print("No metrics found for this document")
						continue
					}

					err = bson.Unmarshal(metrics, &metricsDocument)

					if err != nil {
						log.Print("Error decoding metrics, skipping...",)
						continue
					}

					esDocument.MongoDocument.Metrics = metricsDocument
					err = elastic.InsertDocument(elasticClient, esDocument, "apis")

					if err != nil {
						log.Print("Error inserting document, skipping...",)
						continue
					}
				} else {
					log.Print("No embedding was produced, skipping...")
				}
			} else {
				log.Printf("Already exists, skipping...")
			}

			stop := time.Now().UnixNano()
			prec = stop-start

			current++
		}
	}

	return fn
}
