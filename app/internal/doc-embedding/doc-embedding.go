package doc_embedding


func EmbedFragment(fragments []string, isQuery bool) *Embeddings {
	preprocessed := PreprocessFragment(fragments, isQuery)
	embeddings := Embed(preprocessed)

	return embeddings
}
