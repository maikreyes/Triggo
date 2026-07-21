package config

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Secret     string
	DSN        string
	FrontUrl   string
	AppID      int64
	PrivateKey []byte
}

func NewConfig() *Config {

	appIDStr := os.Getenv("GITHUB_APP_ID")
	appID, err := strconv.ParseInt(appIDStr, 10, 64)
	if err != nil {
		log.Fatal("CRITICAL ERROR: GITHUB_APP_ID is not set or is not a valid number")
	}

	privateKeyB64 := os.Getenv("GITHUB_PRIVATE_KEY_BASE64")
	if privateKeyB64 == "" {
		log.Fatal("CRITICAL ERROR: GITHUB_PRIVATE_KEY_BASE64 is not set")
	}

	privateKey, err := base64.StdEncoding.DecodeString(privateKeyB64)
	if err != nil {
		log.Fatal("CRITICAL ERROR: Decoding of GITHUB_PRIVATE_KEY_BASE64 failed. Verify that you copied the string correctly..")
	}

	return &Config{
		Secret:     os.Getenv("GITHUB_WEBHOOK_SECRET"),
		DSN:        os.Getenv("POSTGRES_URL_NON_POOLING"),
		FrontUrl:   os.Getenv("FRONT_URL"),
		AppID:      appID,
		PrivateKey: privateKey,
	}
}
