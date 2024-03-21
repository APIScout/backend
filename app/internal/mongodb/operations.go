package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SearchDocument - search a document in a collection based on a query. A mongodb database, a query and a collection
// need to be passed to the function.
func SearchDocument(database *mongo.Database, query bson.M, collection string) (bson.Raw, error) {
	coll := database.Collection(collection)
	res, err := coll.FindOne(context.Background(), query).Raw()

	return res, err
}

func InsertDocuments(database *mongo.Database, documents []interface{}, collection string) (*mongo.InsertManyResult, error) {
	coll := database.Collection(collection)
	ids, err := coll.InsertMany(context.Background(), documents)

	return ids, err
}
