package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func cleanPrivateKey(rawKey string) string {
	lines := strings.Split(rawKey, "\n")
	cleanedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		if idx := strings.Index(line, "|"); idx != -1 {
			cleanedLines = append(cleanedLines, strings.TrimSpace(line[idx+1:]))
		} else {
			cleanedLines = append(cleanedLines, strings.TrimSpace(line))
		}
	}

	return strings.Join(cleanedLines, "\n")
}

func CreateToken(ttl time.Duration, payload any, privateKey string) (string, error) {
	cleanedKey := cleanPrivateKey(privateKey)

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(cleanedKey)

	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}

	fmt.Println("Decoded Private Key: ", string(decodedPrivateKey))

	block, _ := pem.Decode(decodedPrivateKey)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("could not parse private key: %w", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)

	if !ok {
		return "", fmt.Errorf("expected an RSA private key")
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub": payload,
		"exp": now.Add(ttl).Unix(),
		"iat": now.Unix(),
		"nbf": now.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaKey)

	if err != nil {
		return "", fmt.Errorf("could not sign token: %w", err)
	}

	return token, nil
}

func ValidateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key: %w", err)
	}

	var key *rsa.PublicKey
	if block, _ := pem.Decode(decodedPublicKey); block != nil {
		key, err = jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
		if err != nil {
			return nil, fmt.Errorf("could not parse PEM public key: %w", err)
		}
	} else {
		cert, err := x509.ParseCertificate(decodedPublicKey)
		if err != nil {
			return nil, fmt.Errorf("could not parse certificate: %w", err)
		}
		key = cert.PublicKey.(*rsa.PublicKey)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims["sub"], nil
}
