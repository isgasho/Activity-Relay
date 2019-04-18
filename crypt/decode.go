package crypt

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

func DecodePublicKeyFromString(publicKeyPem string) (*rsa.PublicKey, error) {
	byteString := []byte(publicKeyPem)
	pemDecoded, _ := pem.Decode(byteString)

	// Is PCKS#1
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(pemDecoded.Bytes)
	if err == nil {
		return rsaPublicKey, err
	}

	// Is PKIX
	keyInterface, err := x509.ParsePKIXPublicKey(pemDecoded.Bytes)
	switch publicKey := keyInterface.(type) {
	case *rsa.PublicKey:
		return publicKey, err
	case *dsa.PublicKey:
		return nil, errors.New("Unsupported publicKey format : DSA")
	case *ecdsa.PublicKey:
		return nil, errors.New("Unsupported publicKey format : ECDSA")
	default:
		return nil, errors.New("Unsupported publicKey format : Unknown")
	}
}

func DecodePrivateKeyFromPath(path string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	decoded, _ := pem.Decode(file)
	privateKey, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
