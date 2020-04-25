package main

import (
	"fmt"
	"log"
	"net/http"
	"workshop/internal/api/jokes"

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

	j := jokes.NewJokeClient(c.JokeUrl)
	h := handler.NewHandler(j)
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)
	log.Printf("Start server at: %s", addr)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
