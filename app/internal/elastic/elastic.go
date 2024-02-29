package elastic

import (
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)


func Connect(host string, port int, user string, password string) (*elasticsearch.Client, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("https://%s:%d", host, port),
		},
		Username: user,
		Password: password,
		CACert:   GetCertificate(),
	}

	return elasticsearch.NewClient(esConfig)
}

func GetCertificate() []byte {
	pwd, _ := os.Getwd()
	cert, _ := os.ReadFile(pwd + "/ca.crt")

	return cert
}
