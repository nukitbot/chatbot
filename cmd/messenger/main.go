package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"nukitbot.github.io/pkg/common"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var v any
		je := json.NewDecoder(r.Body)
		err := je.Decode(&v)
		if err != nil {
			log.Fatalf("Some problem occured parsing request body: %v\n", err)
		}
		log.Printf("Got request: %v\n", v)

		w.WriteHeader(http.StatusOK)
	})

	r.Get("/webhook", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		log.Println(
			q.Get("hub.mode"),
			q.Get("hub.verify_token"),
			q.Get("hub.challenge"),
		)

		w.WriteHeader(http.StatusOK)
	})

	addr := ":" + common.GetEnv("PORT", "80")
	log.Printf("Starting server at %s\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Failed to start the server: %v\n", err)
	}
}
