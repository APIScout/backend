package doc_embedding

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)


type Embeddings struct {
	Predictions [][]float32
}


func Embed(fragments []string) *Embeddings {
	body, _ := json.Marshal(map[string][]string{
		"instances": fragments,
	})

	// Call embedding model
	reqBody := bytes.NewBuffer(body)
	res, err := http.Post(
		"http://127.0.0.1:8501/v1/models/universal_encoder:predict",
		"application/json",
		reqBody,
	)

	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	// Decode the JSON body containing the embeddings
	embeddings := new(Embeddings)
	err = json.NewDecoder(res.Body).Decode(embeddings)

	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	return embeddings
}
