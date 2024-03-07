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

func SearchDocument(client *elasticsearch.Client, query string, index string) {
	request := esapi.SearchRequest{
		Index: []string{index},
		Body: strings.NewReader(query),
	}

	response, err := request.Do(context.TODO(), client)

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] %s", response.String())
	}
}

func DeleteDocument(client *elasticsearch.Client, id string, index string) {
	response, err := client.Delete(index, id)

	if err != nil {
		return
	}

	if os.Getenv("GIN_MODE") == "debug" {
		log.Printf("[ELASTIC-debug] %s", response.String())
	}
}
