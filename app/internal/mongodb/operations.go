package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// RetrieveDocument - search a document in a collection based on a query. A mongodb database, a query and a collection
// need to be passed to the function.
func RetrieveDocument(database *mongo.Database, query bson.M, collection string) (bson.Raw, error) {
	coll := database.Collection(collection)
	res, err := coll.FindOne(context.Background(), query).Raw()

	return res, err
}

// RetrieveDocuments - search documents in a collection based on a query. A mongodb database, a query and a collection
// need to be passed to the function.
func RetrieveDocuments(database *mongo.Database, query bson.D, collection string) (*mongo.Cursor, error) {
	coll := database.Collection(collection)
	res, err := coll.Find(context.Background(), query)

	return res, err
}

// InsertDocuments - insert documents in a collection based on a query. A mongodb database, a document and a collection
// need to be passed to the function.
func InsertDocuments(database *mongo.Database, documents []interface{}, collection string) (*mongo.InsertManyResult, error) {
	coll := database.Collection(collection)
	ids, err := coll.InsertMany(context.Background(), documents)

	return ids, err
}
