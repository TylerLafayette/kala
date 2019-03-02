package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/TylerLafayette/kala/backend/database"
	"github.com/TylerLafayette/kala/backend/models"
)

// LoadTransactionsIntoDb = /api/loadTransactionsIntoDb
func LoadTransactionsIntoDb(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./routes/output.json")
	if err != nil {
		log.Fatal(err)
	}
	var txns []models.Transaction
	err = json.Unmarshal(data, &txns)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range txns {
		log.Println(e.Description)
		var id int64
		err := database.DB.Conn.QueryRow(`INSERT INTO transactions (description, price, final_balance) VALUES($1, $2, $3) RETURNING transaction_id`, e.Description, -1*e.Price, e.FinalBalance).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	render.JSON(w, r, txns)
}
