package elastic

import (
	"backend/app/internal/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseEmbedding - converts the SyncDocument returned by the mongo client, as well as the embeddings returned by the
// Universal Sentence Encoder model, into an EsDocument. A mongo document and an embedding need to be passed to the
// function.
func ParseEmbedding(document *structs.SyncDocument, embeddings *structs.EmbeddingResponse) *structs.EsDocument {
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
