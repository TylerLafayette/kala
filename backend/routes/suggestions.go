package routes

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/utils"
	"github.com/go-chi/render"
)

// Suggestions = /api/suggestions
func Suggestions(w http.ResponseWriter, r *http.Request) {
	transactions, err := database.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}
	suggestions, err := utils.GenerateSuggestions(2.0, 3.0, transactions)
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, suggestions)
}
