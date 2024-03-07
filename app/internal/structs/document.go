package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {}

type SyncDocument struct {
	Document
	Id string `bson:"_id"`
	Collection string `bson:"coll"`
	Api string
}

type EsDocument struct {
	Document
	MongoId primitive.ObjectID `json:"mongo_id"`
	Collection string `json:"mongo_collection"`
	Embedding []float32 `json:"embedding"`
}
