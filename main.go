package main

import (
	"apiSecurity/server"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
