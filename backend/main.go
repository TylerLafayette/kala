package main

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/routes"
	"github.com/go-chi/chi"
)

func main() {
	// Connect to the database.
	database.Connect()

	// Create a router and declare routes.
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/overview", routes.Overview)
		r.Get("/transactions", routes.Transactions)
		r.Get("/suggestions", routes.Suggestions)
		r.Get("/loadTransactionsIntoDb", routes.LoadTransactionsIntoDb)
	})

	// Start serving HTTP.
	log.Fatal(http.ListenAndServe(":7800", r))
}
