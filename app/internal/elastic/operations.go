package elastic

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"backend/app/internal/models"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/goccy/go-json"
)

// InsertDocument - send and save a document in an elasticsearch index. An elasticsearch client, document and index need
// to be passed to the function.
func InsertDocument(client *elasticsearch.Client, document models.EsRequest, index string) error {
	jsonDocument, err := json.Marshal(document)

	if err != nil {
		return err
	}

	res, err := client.Index(index, bytes.NewReader(jsonDocument))

	if res != nil && res.StatusCode == http.StatusCreated {
		if os.Getenv("GIN_MODE") == "debug" {
			log.Printf("[ELASTIC-debug] Indexing: %s", res.Status())
		}
	} else {
		return errors.New("could not insert document")
	}

	return err
}

// SearchDocument - search a document in an index based on a query. An elasticsearch client, a query and an index need
// to be passed to the function.
func SearchDocument(client *elasticsearch.Client, query string, index string) (*models.EsSearchResponse, error) {
	var response models.EsSearchResponse

	res, err := client.Search(
		client.Search.WithIndex(index),
		client.Search.WithBody(strings.NewReader(query)),
		client.Search.WithContext(context.TODO()),
	)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	// Unmarshal the byte array into the response
	err = json.Unmarshal([]byte(out.String()), &response)

	return &response, err
}

func GetDocumentCount(client *elasticsearch.Client, index string) (int64, error) {
	var docsCount struct{Count int64 `json:"count"`}

	res, err := client.Count(
		client.Count.WithIndex(index),
		client.Count.WithContext(context.TODO()),
	)

	if err != nil {
		return 0, err
	}

	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(res.Body, b1)
	_, err = out.ReadFrom(tr)

	if err != nil {
		return 0, err
	}

	err = json.Unmarshal([]byte(out.String()), &docsCount)

	return docsCount.Count, err
}

// DeleteDocument - delete a document in an index based on its id. An elasticsearch client, an id and an index need to
// be passed to the function.
func DeleteDocument(client *elasticsearch.Client, id string, index string) error {
	res, err := client.Delete(index, id)

	if err == nil {
		if os.Getenv("GIN_MODE") == "debug" {
			log.Printf("[ELASTIC-debug] Delete: %s", res.Status())
		}
	}

	return err
}
