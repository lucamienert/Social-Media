package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(ttl time.Duration, payload any, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)

	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}

	key, err := x509.ParsePKCS8PrivateKey(decodedPrivateKey)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(decodedPrivateKey)

		if err != nil {
			return "", fmt.Errorf("could not parse private key: %w", err)
		}
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
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
