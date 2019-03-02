package main

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/routes"
	"github.com/go-chi/chi"
)

func main() {
	database.Connect()
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/overview", routes.Overview)
	})
	log.Fatal(http.ListenAndServe(":7800", r))
}
