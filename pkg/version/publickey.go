package version

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func PublicKey() (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(`
-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAFTyeOLfBS1ty0NAYOzmytiraWLn/
2D9XMvPkgo9woP9GB3qeY1HcO38SCUUzdZug9ZLk9mYam7IbGdbPMbSXkOYBxXvG
rbOcyQmRniXyYcrIrBU20qt92ozGXj7y2YMpyPOS5wBJLO/gORnKN8OCEKvAXXz0
C2lUZ6TW2UVhmDeGrbU=
-----END PUBLIC KEY-----`))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Failed to parse public key: " + err.Error())
	}
	pubKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("PublicKey is not ECDSA")
	}
	return pubKey, nil
}
