package main

import (
	"github.com/emanuel3k/playlist-transfer/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	if err := config.InitDB(); err != nil {
		log.Fatal(err)
	}

	if err := config.InitHTTPServer(); err != nil {
		log.Fatal(err)
	}
}
