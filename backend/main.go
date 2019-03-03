package main

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	// Connect to the database.
	database.Connect()

	// Create a router and declare routes.
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Route("/api", func(r chi.Router) {
		r.Get("/overview", routes.Overview)
		r.Get("/transactions", routes.Transactions)
		r.Get("/suggestions", routes.Suggestions)
		r.Post("/insert", routes.Insert)
		r.Get("/loadTransactionsIntoDb", routes.LoadTransactionsIntoDb)
	})

	// Start serving HTTP.
	log.Fatal(http.ListenAndServe(":7800", r))
}
