package models

// EmbeddingRequest - structure of the request to be sent to the embedding server.
type EmbeddingRequest struct {
	Fragment string `json:"fragment"`
}

// EmbeddingResponse - structure of the response sent back by the embedding server.
type EmbeddingResponse struct {
	Predictions [][]float32
}
