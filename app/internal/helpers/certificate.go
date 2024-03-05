package helpers

import (
	"os"
)

func GetCertificate() []byte {
	pwd, _ := os.Getwd()
	cert, _ := os.ReadFile(pwd + "/ca.crt")

	return cert
}
