package utils

import (
	"fmt"
	"time"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"
)

func HashPass(password string) string {

	saltPass := viper.GetString("PASSWORD_SALT")

	hashedPassword := argon2.IDKey([]byte(password), []byte(saltPass), 1, 64*1024, 4, 32)

	hashedPasswordHex := fmt.Sprintf("%x", hashedPassword)

	return string(hashedPasswordHex)
}

func PassValidation(password string, hashPassword string) bool {

	requestPassword := HashPass(password)

	if requestPassword == hashPassword {
		return true
	} else {
		return false
	}

}

func GenerateKSUID() (string, error) {

	generatedId, err := ksuid.NewRandomWithTime(time.Now())
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	return generatedId.String(), nil

}
