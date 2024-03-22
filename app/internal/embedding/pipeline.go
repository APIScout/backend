package embedding

import (
	"backend/app/internal/models"
)

// PerformPipeline - fragments are preprocessed and embeddings are generated and returned. An array of fragments
// (string) and a boolean indicating if the fragments are queries or not need to be passed to the function.
func PerformPipeline(fragments []string, isQuery bool) (*models.EmbeddingResponse, error) {
	preprocessed := PreprocessFragment(fragments, isQuery)
	embeddings, err := Embed(preprocessed)

	return embeddings, err
}
