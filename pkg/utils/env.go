package utils

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func GetEnv(k string) (value string) {
	value = os.Getenv(k)
	if value == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	return
}

func GetBoolEnv(k string) (b bool) {
	value := os.Getenv(k)
	if value == "" {
		log.Panicln("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		log.Panicln("ENV err: [" + k + "]\n" + err.Error())
	}
	return
}
