package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDbUrl() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", GetEnv("DB_USER"), GetEnv("DB_PW"), GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_NAME"))
}

func GetEnv(key string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func CheckError(err error, args ...interface {
}) {
	if err != nil {
		log.Fatalln(err, args)
	}
}
