package embedding

func PerformPipeline(fragments []string, isQuery bool) *Embeddings {
	preprocessed := PreprocessFragment(fragments, isQuery)
	embeddings := Embed(preprocessed)

	return embeddings
}
