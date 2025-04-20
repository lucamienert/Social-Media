package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKey(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func savePEMKey(filename string, keyBytes []byte, keyType string) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	return pem.Encode(file, &pem.Block{
		Type:  keyType,
		Bytes: keyBytes,
	})
}

func savePrivateKey(privateKey *rsa.PrivateKey) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	return savePEMKey("private_key.pem", keyBytes, "RSA PRIVATE KEY")
}

func savePublicKey(publicKey *rsa.PublicKey) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return err
	}

	return savePEMKey("public_key.pem", pubASN1, "PUBLIC KEY")
}

func main() {
	bits := 2048
	privateKey, err := generateRSAKey(bits)

	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}

	if err := savePrivateKey(privateKey); err != nil {
		fmt.Println("Error saving private key:", err)
		return
	}

	fmt.Println("Private key saved as private_key.pem")

	if err := savePublicKey(&privateKey.PublicKey); err != nil {
		fmt.Println("Error saving public key:", err)
		return
	}

	fmt.Println("Public key saved as public_key.pem")
}
