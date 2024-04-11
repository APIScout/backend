package models

// EsRequest - structure of an elasticsearch document to be sent to the elasticsearch client.
type EsRequest struct {
	MongoDocument MongoDocument `json:"metadata"`
	Embedding     []float32     `json:"embedding"`
}

// EsSearchResponse - structure of the response sent by the elasticsearch client
type EsSearchResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

// Hit - structure of an elasticsearch document
type Hit struct {
	Id       string `json:"_id"`
	Index    string `json:"_index"`
	Document struct {
		Metadata struct {
			MongoId string `json:"mongo-id"`
		} `json:"metadata"`
		Embedding []float32 `json:"embedding"`
	} `json:"_source"`
}
