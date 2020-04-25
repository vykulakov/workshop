package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	c := &config.Server{}
	err := cleanenv.ReadConfig("config.yml", c)
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	h := handler.NewHandler()

	r.Get("/hello", h.Hello)

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	log.Printf("Start server at: %s", addr)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
