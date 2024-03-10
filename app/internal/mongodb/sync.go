package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"backend/app/internal/models"

	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WatchDatabase - watch for a certain type of events in the mongodb `apis` database. The supported events are
// document insertion and document deletion.
func WatchDatabase(client *mongo.Client, esClient *elasticsearch.Client, operation string) {
	db := client.Database("apis")
	match := bson.D{{"$match", bson.D{{"operationType", operation}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(5 * time.Second)

	stream, err := db.Watch(context.TODO(), mongo.Pipeline{match}, opts)

	if err != nil {
		log.Fatal(err)
	}

	switch operation {
	case "insert":
		InsertDocuments(esClient, stream, db)
	case "delete":
		DeleteDocuments(esClient, stream, db)
	}
}

// InsertDocuments - every time a document is inserted in the database, it will be taken from mongodb, embedded, and
// saved in the respective elasticsearch index.
func InsertDocuments(esClient *elasticsearch.Client, stream *mongo.ChangeStream, database *mongo.Database) {
	for stream.Next(context.TODO()) {
		var document models.SyncDocument
		// Retrieve the mongo document id
		err := stream.Current.Lookup("documentKey").Unmarshal(&document)

		if err != nil {
			panic(err)
		}

		// Retrieve the mongo document collection
		err = stream.Current.Lookup("ns").Unmarshal(&document)

		if err != nil {
			panic(err)
		}

		// Create ObjectId for mongodb query
		docId, err := primitive.ObjectIDFromHex(document.Id)

		if err != nil {
			panic(err)
		}

		query := bson.M{"_id": docId}
		document.Api = SearchDocument(database, query, document.Collection)

		embeddings := embedding.PerformPipeline([]string{document.Api}, false)
		esDocument := elastic.ParseEmbedding(&document, embeddings)
		// TODO: Change index name
		elastic.SendDocument(esClient, esDocument, "test")
	}
}

// DeleteDocuments - every time a documents is deleted from the database, it will be searched in the elasticsearch
// index (based on the mongodb ObjectId) and deleted from the elasticsearch database.
func DeleteDocuments(esClient *elasticsearch.Client, stream *mongo.ChangeStream, database *mongo.Database) {
	for stream.Next(context.TODO()) {
		var specification models.SyncDocument
		err := stream.Current.Lookup("documentKey").Unmarshal(&specification)

		if err != nil {
			panic(err)
		}

		query := fmt.Sprintf(`{"query": {"match": {"mongo_id": "%s"}}}`, specification.Id)
		// TODO: Change index name
		elastic.SearchDocument(esClient, query, "test")
		//elastic.DeleteDocument(esClient, esDocument.)
	}
}
