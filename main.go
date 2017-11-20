package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kamilsk/form-api/server"
	"github.com/kamilsk/form-api/server/router/chi"
	_ "github.com/lib/pq"
)

func main() {
	addr := addr()
	log.Println("starting server at", addr)
	log.Fatal(http.ListenAndServe(addr, chi.NewRouter(
		server.New())))
}

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return os.Getenv("BIND") + ":" + port
}
