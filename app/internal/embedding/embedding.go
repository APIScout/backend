package embedding

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"backend/app/internal/structs"
)

type Embeddings = structs.EmbeddingResponse

// Embed use the Universal Sentence Encoder model to transform the array of fragments (string) into an array of
// embeddings (512-dimension float32 embedding). A list of embeddings needs to be passed to the function.
func Embed(fragments []string) *Embeddings {
	body, _ := json.Marshal(map[string][]string{
		"instances": fragments,
	})

	// Call embedding model
	reqBody := bytes.NewBuffer(body)
	res, err := http.Post(
		"http://"+os.Getenv("MODELS_HOST")+":8501/v1/models/universal-encoder:predict",
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
