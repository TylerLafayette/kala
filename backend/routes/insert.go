package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/TylerLafayette/kala/backend/database"
)

// Insert = /api/insert
func Insert(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	var id int64
	err = database.DB.Conn.QueryRow(`INSERT INTO transactions (description, price, final_balance) VALUES($1, $2, $3) RETURNING transaction_id`, m["description"].(string), m["price"].(float64), m["final_balance"].(float64)).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	m["transaction_id"] = id

	i := (*spendingStat)[0] + float32(m["price"].(float64))

	spendingStat = &map[int]float32{
		0: i,
	}

	render.JSON(w, r, m)
}
