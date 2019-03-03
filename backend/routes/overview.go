package routes

import (
	"log"
	"net/http"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/models"
	"github.com/TylerLafayette/kala/backend/utils"

	"github.com/go-chi/render"
)

var spendingStat = &map[int]float32{
	0: 500,
}

// Overview = /api/overview
func Overview(w http.ResponseWriter, r *http.Request) {
	transactions, err := database.GetTransactions()
	if err != nil {
		log.Fatal(err)
	}

	suggestions, err := utils.GenerateSuggestions(300, 200, transactions)
	if err != nil {
		log.Fatal(err)
	}

	o := models.Overview{
		Username:       "Lucas",
		CurrentBalance: 300,
		AmountSpent:    (*spendingStat)[0],
		Budget:         2400,
		Transactions:   transactions,
		Suggestions:    suggestions,
	}

	render.JSON(w, r, o)
}
