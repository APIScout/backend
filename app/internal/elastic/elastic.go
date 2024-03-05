package elastic

import (
	"fmt"

	"backend/app/internal/helpers"
	"backend/app/internal/structs"

	"github.com/elastic/go-elasticsearch/v8"
)

func Connect(config structs.Elastic) *elasticsearch.Client {
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

	return client
}
