package main

import (
	"apiSecurity/server"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
