package cert

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func PemTox509(input []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(input)
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}
	return x509.ParseCertificate(block.Bytes)
}
