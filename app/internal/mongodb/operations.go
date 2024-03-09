package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SearchDocument - search a document in a collection based on a query. A mongodb database, a query and a collection
// need to be passed to the function.
func SearchDocument(database *mongo.Database, query bson.M, collection string) string {
	coll := database.Collection(collection)
	res, err := coll.FindOne(context.Background(), query).Raw()

	if err != nil {
		panic(err)
	}

	return string(res)
}
