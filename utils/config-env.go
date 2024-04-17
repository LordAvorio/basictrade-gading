package utils

import "github.com/joho/godotenv"

func ReadConfigEnvironment() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}