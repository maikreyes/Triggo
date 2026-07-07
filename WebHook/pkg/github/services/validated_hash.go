package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strings"
)

func (s *Services) ValidatedHash(signature string, payload []byte) error {

	log.Println(signature)

	hash := strings.TrimPrefix(signature, "sha256=")

	log.Println(hash)

	if hash == "" {
		return errors.New("Signature invalid")
	}

	secret := s.Config.Secret

	log.Printf("secret:%s\n", secret)

	decoded, err := hex.DecodeString(hash)

	if err != nil {
		return errors.New("Error to decode hash")
	}

	mac := hmac.New(sha256.New, []byte(secret))

	mac.Write(payload)

	calculatedHash := mac.Sum(nil)

	log.Printf("decoded:%s\n", string(decoded))
	log.Printf("calculatedHash:%s\n", string(calculatedHash))

	if !hmac.Equal(decoded, calculatedHash) {
		return errors.New("The signature is invalid")
	}

	return nil
}
