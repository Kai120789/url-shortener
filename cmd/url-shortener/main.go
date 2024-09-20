package main

import (
	"fmt"
	"log"
	"url-shortener/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../local.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: init server

}
