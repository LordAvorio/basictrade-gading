package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"
)

func HashPass(password string) string {

	saltPass := viper.GetString("PASSWORD_SALT")

	hashedPassword := argon2.IDKey([]byte(password), []byte(saltPass), 1, 64*1024, 4, 32)

	hashedPasswordHex := fmt.Sprintf("%x", hashedPassword)

	return string(hashedPasswordHex)
}
