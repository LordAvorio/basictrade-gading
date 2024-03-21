package helpers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateToken(uuid string, email string) (string, error) {

	claims := jwt.MapClaims{
		"uuid":    uuid,
		"email":   email,
		"expired": time.Now().Add(time.Minute * 30).Unix(),
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signedToken, err := parseToken.SignedString(privateKey)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return signedToken, nil

}
