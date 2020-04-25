package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/api/jokes"
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
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit

		log.Print("Received the stop signal")

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		done <- srv.Shutdown(ctx)
	}()

	log.Printf("Start server at: %s", addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Server stopped")
}
