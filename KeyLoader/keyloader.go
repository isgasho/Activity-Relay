package keyloader

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadPrivateKeyRSAfromPath(path string) (*rsa.PrivateKey, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	decoded, _ := pem.Decode(file)
	priv, err := x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func ReadPublicKeyRSAfromString(pemString string) (*rsa.PublicKey, error) {
	pemByte := []byte(pemString)
	decoded, _ := pem.Decode(pemByte)
	keyInterface, err := x509.ParsePKIXPublicKey(decoded.Bytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	pub := keyInterface.(*rsa.PublicKey)
	return pub, nil
}

func GeneratePublicKeyPEMString(publicKey *rsa.PublicKey) string {
	publicKeyByte := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyByte,
		},
	)
	return string(publicKeyPem)
}
