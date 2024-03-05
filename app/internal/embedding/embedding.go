package embedding

import (
	"backend/app/internal/structs"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Embeddings = structs.Embeddings

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
