package helpers

import (
	"fmt"
	"os"
)

// GetCertificate - retrieve the elasticsearch cluster certificate.
func GetCertificate() []byte {
	pwd, _ := os.Getwd()
	cert, _ := os.ReadFile(pwd + fmt.Sprintf("/%s-ca.crt", os.Getenv("GIN_MODE")))

	return cert
}
