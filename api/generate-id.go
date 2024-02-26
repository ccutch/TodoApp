package api

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateID(len int) string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		log.Fatal("failed to generate random secret")
	}
	return base64.URLEncoding.EncodeToString(b)
}
