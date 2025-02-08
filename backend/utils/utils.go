package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// Gera um token aleatório com o comprimento especificado
func GenerateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
