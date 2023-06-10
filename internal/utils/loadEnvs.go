package utils

import "github.com/joho/godotenv"

func LoadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
