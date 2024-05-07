package embedding

import (
	"backend/app/internal/models"
	"errors"
)

// PerformPipeline - fragments are preprocessed and embeddings are generated and returned. An array of fragments
// (string) and a boolean indicating if the fragments are queries or not need to be passed to the function.
func PerformPipeline(fragments []string, isQuery bool) (*models.EmbeddingResponse, []int, error) {
	var preprocessedLengths []int
	preprocessed := PreprocessFragment(fragments, isQuery)

	if preprocessed == nil {
		return nil, []int{}, errors.New("no fragments were generated")
	}

	embeddings, err := Embed(preprocessed)

	if err != nil {
		return nil, []int{}, err
	}

	for _, preprocessedItem := range preprocessed {
		preprocessedLengths = append(preprocessedLengths, len(preprocessedItem))
	}

	return embeddings, preprocessedLengths, nil
}
