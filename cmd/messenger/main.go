package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		w.Write([]byte(q.Get("hub.challenge")))
	})

	addr := ":" + os.Getenv("PORT")
	log.Printf("Starting server at %s\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Failed to start the server: %v\n", err)
	}
}
