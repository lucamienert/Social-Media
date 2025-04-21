package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func encodePrivateKeyToBase64(privateKey *rsa.PrivateKey) (string, error) {
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPem := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes})
	return base64.StdEncoding.EncodeToString(privPem), nil
}

func encodePublicKeyToBase64(publicKey *rsa.PublicKey) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	pubPem := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
	return base64.StdEncoding.EncodeToString(pubPem), nil
}

func main() {
	accessPrivateKey, accessPublicKey, err := generateRSAKeyPair(2048)

	if err != nil {
		log.Fatalf("Error generating access key pair: %v", err)
	}

	refreshPrivateKey, refreshPublicKey, err := generateRSAKeyPair(2048)

	if err != nil {
		log.Fatalf("Error generating refresh key pair: %v", err)
	}

	accessPrivateKeyBase64, err := encodePrivateKeyToBase64(accessPrivateKey)

	if err != nil {
		log.Fatalf("Error encoding access private key to base64: %v", err)
	}
	accessPublicKeyBase64, err := encodePublicKeyToBase64(accessPublicKey)

	if err != nil {
		log.Fatalf("Error encoding access public key to base64: %v", err)
	}

	refreshPrivateKeyBase64, err := encodePrivateKeyToBase64(refreshPrivateKey)

	if err != nil {
		log.Fatalf("Error encoding refresh private key to base64: %v", err)
	}

	refreshPublicKeyBase64, err := encodePublicKeyToBase64(refreshPublicKey)

	if err != nil {
		log.Fatalf("Error encoding refresh public key to base64: %v", err)
	}

	fmt.Println("ACCESS_TOKEN_PRIVATE_KEY (Base64):")
	fmt.Println(accessPrivateKeyBase64)

	fmt.Println("ACCESS_TOKEN_PUBLIC_KEY (Base64):")
	fmt.Println(accessPublicKeyBase64)

	fmt.Println("REFRESH_TOKEN_PRIVATE_KEY (Base64):")
	fmt.Println(refreshPrivateKeyBase64)

	fmt.Println("REFRESH_TOKEN_PUBLIC_KEY (Base64):")
	fmt.Println(refreshPublicKeyBase64)

	err = os.WriteFile("access_private_key_base64.txt", []byte(accessPrivateKeyBase64), 0644)

	if err != nil {
		log.Fatalf("Error writing access private key to file: %v", err)
	}

	err = os.WriteFile("access_public_key_base64.txt", []byte(accessPublicKeyBase64), 0644)

	if err != nil {
		log.Fatalf("Error writing access public key to file: %v", err)
	}

	err = os.WriteFile("refresh_private_key_base64.txt", []byte(refreshPrivateKeyBase64), 0644)

	if err != nil {
		log.Fatalf("Error writing refresh private key to file: %v", err)
	}

	err = os.WriteFile("refresh_public_key_base64.txt", []byte(refreshPublicKeyBase64), 0644)

	if err != nil {
		log.Fatalf("Error writing refresh public key to file: %v", err)
	}
}
