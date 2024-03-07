package mongodb

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"time"

	"backend/app/internal/elastic"
	"backend/app/internal/embedding"
	"backend/app/internal/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func InsertDocuments(esClient *elasticsearch.Client, stream *mongo.ChangeStream, database *mongo.Database) {
	for stream.Next(context.TODO()) {
		document := RetrieveDocument(stream, database)
		embeddings := embedding.PerformPipeline([]string{document.Api}, false)
		esDocument := elastic.ParseEmbedding(document, embeddings)
		// TODO: Change index name
		elastic.SendDocument(esClient, esDocument, "test")
	}
}

func DeleteDocuments(esClient *elasticsearch.Client, stream *mongo.ChangeStream, database *mongo.Database) {
	for stream.Next(context.TODO()) {
		document := RetrieveDocument(stream, database)
		query := fmt.Sprintf(`{"query": {"match": {"mongo_id": "%s"}}}`, document.Id)

		// TODO: Change index name
		elastic.SearchDocument(esClient, query, "test")
	}
}

func RetrieveDocument(stream *mongo.ChangeStream, db *mongo.Database) *structs.SyncDocument {
	var specification structs.SyncDocument
	err := stream.Current.Lookup("documentKey").Unmarshal(&specification)

	if err != nil {
		panic(err)
	}

	err = stream.Current.Lookup("ns").Unmarshal(&specification)

	if err != nil {
		panic(err)
	}

	id, err := primitive.ObjectIDFromHex(specification.Id)

	if err != nil {
		panic(err)
	}

	coll := db.Collection(specification.Collection)
	res, err := coll.FindOne(context.Background(), bson.M{"_id": id}).Raw()

	if err != nil {
		panic(err)
	}

	specification.Api = fmt.Sprint(res)

	return &specification
}
