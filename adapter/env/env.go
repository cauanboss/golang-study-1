package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func GetHttpPort() string {
	return os.Getenv("http_port")
}
