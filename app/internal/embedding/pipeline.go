package embedding

// PerformPipeline - fragments are preprocessed and embeddings are generated and returned. An array of fragments
// (string) and a boolean indicating if the fragments are queries or not need to be passed to the function.
func PerformPipeline(fragments []string, isQuery bool) *Embeddings {
	preprocessed := PreprocessFragment(fragments, isQuery)
	embeddings := Embed(preprocessed)

	return embeddings
}
