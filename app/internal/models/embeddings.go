package models

var PossibleFilters = []string{"metadata", "specification"}

// EmbeddingRequest - structure of the request to be sent to the embedding server.
type EmbeddingRequest struct {
	Fragment string `json:"fragment,omitempty"`
	DSL      string `json:"filters,omitempty"`
	Fields   []string `json:"fields,omitempty"`
}

// EmbeddingResponse - structure of the response sent back by the embedding server.
type EmbeddingResponse struct {
	Predictions [][]float32
}
