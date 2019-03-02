package routes

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/go-chi/render"
)

// Transactions = /api/transactions
func Transactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := database.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}

	render.JSON(w, r, transactions)
}
