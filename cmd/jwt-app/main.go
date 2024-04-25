package main

import (
	"log"
	"net/http"

	"github.com/MelihcanSrky/GolangJWT/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	mux.Post("/login", handlers.LoginHandler)
	mux.Get("/protected", handlers.ProtectedHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
