package embedding

import (
	"backend/app/internal/models"
	"errors"
)

// PerformPipeline - fragments are preprocessed and embeddings are generated and returned. An array of fragments
// (string) and a boolean indicating if the fragments are queries or not need to be passed to the function.
func PerformPipeline(fragments []string, isQuery bool) (*models.EmbeddingResponse, int, error) {
	preprocessed := PreprocessFragment(fragments, isQuery)

	if preprocessed == nil {
		return nil, 0, errors.New("no fragments were generated")
	}

	embeddings, err := Embed(preprocessed)

	if err != nil {
		return nil, 0, err
	}

	return embeddings, len(preprocessed[0]), nil
}
