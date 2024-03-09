package elastic

import (
	"bytes"
	"context"
	"log"
	"os"
	"strings"

	"backend/app/internal/structs"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/goccy/go-json"
)

// SendDocument - send and save a document in an elasticsearch index. An elasticsearch client, document and index need
// to be passed to the function.
func SendDocument(client *elasticsearch.Client, document *structs.EsDocument, index string) {
	jsonDocument, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}

	response, err := client.Index(index, bytes.NewReader(jsonDocument))

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] %s", response.String())
	}
}

// SearchDocument - search a document in an index based on a query. An elasticsearch client, a query and an index need
// to be passed to the function.
func SearchDocument(client *elasticsearch.Client, query string, index string) *structs.EsDocument {
	var response structs.EsDocument
	request := esapi.SearchRequest{
		Index: []string{index},
		Body:  strings.NewReader(query),
	}

	res, err := request.Do(context.TODO(), client)

	if err != nil {
		panic(err)
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] %s", res.Status())
	}

	//bson.NewDecoder(res)
	log.Print(response)

	if err != nil {
		panic(err)
	}

	return &response
}

// DeleteDocument - delete a document in an index based on its id. An elasticsearch client, an id and an index need to
// be passed to the function.
func DeleteDocument(client *elasticsearch.Client, id string, index string) {
	response, err := client.Delete(index, id)

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] %s", response.String())
	}
}
