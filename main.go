package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kamilsk/form-api/server/chi"
)

func main() {
	log.Fatal(http.ListenAndServe(addr(), chi.NewRouter()))
}

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return os.Getenv("BIND") + ":" + port
}
