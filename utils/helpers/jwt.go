package helpers

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateToken(id uint, email string) (string, error) {

	claims := jwt.MapClaims{
		"id":      id,
		"email":   email,
		"expired": time.Now().Add(time.Minute * 30),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SALT")
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return signedToken, nil

}

func VerifyToken(ctx *gin.Context) (interface{}, error) {

	tokenHeader := ctx.Request.Header.Get("Authorization")
	secretKey := os.Getenv("JWT_SALT")

	if bearer := strings.HasPrefix(tokenHeader, "Bearer"); !bearer {
		log.Error().Msg("Bearer is empty")
		return nil, errors.New("login to proceed")
	}

	stringToken := strings.Split(tokenHeader, " ")[1]
	tokenJwt, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error().Msg("Unexpected signing method JWT")
			return nil, errors.New("login to proceed")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	claims, ok := tokenJwt.Claims.(jwt.MapClaims)
	if !ok || !tokenJwt.Valid {
		log.Error().Msg("Token is invalid")
		return nil, errors.New("token is invalid")
	}

	expClaim, exists := claims["expired"]
	if !exists {
		log.Error().Msg("Expire claim is missing")
		return nil, errors.New("expire claim is missing")
	}

	expStr, ok := expClaim.(string)
	if !ok {
		log.Error().Msg("Expire claim is not a valid type")
		return nil, errors.New("expire claim is not a valid type")
	}

	expTime, err := time.Parse(time.RFC3339, expStr)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, errors.New("error parsing expiration time")
	}

	if time.Now().After(expTime) {
		log.Error().Msg("Token is expired")
		return nil, errors.New("token is expired")
	}

	return claims, nil


}
