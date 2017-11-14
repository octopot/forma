package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/{UID}", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("get form schema"))
	})
	r.Post("/{UID}", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("send form data"))
	})
	log.Fatal(http.ListenAndServe(addr(), r))
}

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return os.Getenv("BIND") + ":" + port
}
