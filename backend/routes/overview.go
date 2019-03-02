package routes

import (
	"net/http"

	"github.com/TylerLafayette/kala/backend/models"
	"github.com/go-chi/render"
)

var dummyOverview = models.Overview{
	Username:       "Lucas",
	CurrentBalance: 24.99,
	Budget:         2400,
	Transactions: []models.Transaction{
		models.Transaction{
			Description:  "SQ *COFFEEBAR MENLO PAR 11/30 MOBILE PURCHASE Menlo Park CA",
			Price:        -14.99,
			FinalBalance: 249.32,
		},
	},
	Suggestions: []models.Suggestion{
		models.Suggestion{
			Emoji:                "☕️",
			Header:               "Cut down on the coffee.",
			Description:          "Maybe try buying cheaper instant coffee that delivers the same caffeine boost.",
			AmountSpentThisMonth: 49.0,
			CallToActionType:     "SAVING_SCHEDULE",
			Data: map[string]interface{}{
				"amtMonthly": 5.0,
			},
		},
	},
}

// Overview = /api/overview
func Overview(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, dummyOverview)
}
