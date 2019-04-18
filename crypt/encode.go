package crypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func EncodePKCS1FromPrivateKey(privateKey *rsa.PrivateKey) string {
	privateKeyByte := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyByte,
		},
	)
	return string(privateKeyPem)
}

func EncodePKCS1FromPublicKey(publicKey *rsa.PublicKey) string {
	publicKeyByte := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyByte,
		},
	)
	return string(publicKeyPem)
}

func EncodePKCIXFromPublicKey(publicKey *rsa.PublicKey) string {
	publicKeyByte, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyByte,
		},
	)
	return string(publicKeyPem)
}
