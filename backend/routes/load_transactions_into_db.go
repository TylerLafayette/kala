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
	for i, j := 0, len(txns)-1; i < j; i, j = i+1, j-1 {
		txns[i], txns[j] = txns[j], txns[i]
	}
	err = json.Unmarshal(data, &txns)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = database.DB.Conn.Query("DELETE FROM transactions")

	for _, e := range txns {
		log.Println(e.Description)
		var id int64
		err := database.DB.Conn.QueryRow(`INSERT INTO transactions (description, price, final_balance) VALUES($1, $2, $3) RETURNING transaction_id`, e.Description, -1*e.Price, e.FinalBalance).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}

	spendingStat = &map[int]float32{
		0: 500,
	}

	render.JSON(w, r, txns)
}
