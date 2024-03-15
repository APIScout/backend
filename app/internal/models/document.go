package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SyncDocument - structure of the document returned by the mongodb driver and used in the sync process.
type SyncDocument struct {
	Id         string `bson:"_id"`
	Collection string `bson:"coll"`
	Api        bson.Raw
}

// EsDocument - structure of an elasticsearch document to be sent to the elasticsearch client.
type EsDocument struct {
	MongoId    primitive.ObjectID `json:"mongo_id"`
	Collection string             `json:"mongo_collection"`
	Embedding  []float32          `json:"embedding"`
}
