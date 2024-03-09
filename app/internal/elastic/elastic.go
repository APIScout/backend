package elastic

import (
	"fmt"
	"log"

	"backend/app/internal/helpers"
	"backend/app/internal/structs"

	"github.com/elastic/go-elasticsearch/v8"
)

// Connect - used to connect to the elasticsearch database. It will return an elasticsearch client that can be used to
// perform queries on the database.
func Connect(config structs.ElasticConfig) *elasticsearch.Client {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("%s://%s:%d", config.Protocol, config.Host, config.Port),
		},
		Username: config.User,
		Password: config.Password,
		CACert:   helpers.GetCertificate(),
	}

	client, err := elasticsearch.NewClient(esConfig)

	if err != nil {
		panic(err)
	}
	log.Print("Connected to ElasticSearch")

	return client
}
