package helpers

import (
	"os"
)

// GetCertificate - retrieve the elasticsearch cluster certificate.
func GetCertificate() []byte {
	pwd, _ := os.Getwd()
	cert, _ := os.ReadFile(pwd + "/local-ca.crt")

	return cert
}
