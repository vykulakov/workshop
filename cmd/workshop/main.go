package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"workshop/internal/handler"
)

func main() {
	r := chi.NewRouter()
	h := handler.NewHandler()

	r.Get("/hello", h.Hello)

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
