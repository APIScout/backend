package elastic

import (
	"backend/app/internal/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseEmbedding(document *structs.SyncDocument, embeddings *structs.Embeddings) *structs.EsDocument {
	var err error
	var esDocument structs.EsDocument

	esDocument.MongoId, err = primitive.ObjectIDFromHex(document.Id)

	if err != nil {
		panic(err)
	}

	esDocument.Collection = document.Collection
	esDocument.Embedding = embeddings.Predictions[0]

	return &esDocument
}
