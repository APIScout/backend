package elastic

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"strings"

	"backend/app/internal/models"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goccy/go-json"
)

// SendDocument - send and save a document in an elasticsearch index. An elasticsearch client, document and index need
// to be passed to the function.
func SendDocument(client *elasticsearch.Client, document *models.EsDocument, index string) {
	jsonDocument, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}

	res, err := client.Index(index, bytes.NewReader(jsonDocument))

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] Indexing: %s", res.Status())
	}
}

// SearchDocument - search a document in an index based on a query. An elasticsearch client, a query and an index need
// to be passed to the function.
func SearchDocument(client *elasticsearch.Client, query string, index string) *models.EsSearchResponse {
	var response models.EsSearchResponse

	res, err := client.Search(
		client.Search.WithIndex(index),
		client.Search.WithBody(strings.NewReader(query)),
		client.Search.WithContext(context.TODO()),
	)

	if err != nil {
		panic(err)
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] Search: %s", res.Status())
	}

	// Read the body of the elasticsearch response
	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(res.Body, b1)
	_, err = out.ReadFrom(tr)

	if err != nil {
		panic(err)
	}

	// Unmarshal the byte array into the response
	err = json.Unmarshal([]byte(out.String()), &response)

	if err != nil {
		panic(err)
	}

	return &response
}

// DeleteDocument - delete a document in an index based on its id. An elasticsearch client, an id and an index need to
// be passed to the function.
func DeleteDocument(client *elasticsearch.Client, id string, index string) {
	res, err := client.Delete(index, id)

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] Delete: %s", res.Status())
	}
}
