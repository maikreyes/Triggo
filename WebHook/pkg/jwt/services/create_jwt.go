package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Services) CreateJWT() (string, error) {

	var JWTstr string

	key, err := jwt.ParseRSAPrivateKeyFromPEM(s.Config.PrivateKey)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iat": time.Now().Unix(),
			"iss": strconv.FormatInt(s.Config.AppID, 10),
			"exp": time.Now().Add(time.Minute * 10).Unix(),
		})

	JWTstr, err = token.SignedString(key)

	if err != nil {
		return "", err
	}

	return JWTstr, nil

}
