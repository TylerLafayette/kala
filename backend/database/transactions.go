package database

import (
	"strings"

	"github.com/TylerLafayette/kala/backend/models"
)

// GetTransactions returns all transactions.
func GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := "SELECT description, price, final_balance FROM transactions ORDER BY transaction_id DESC"
	rows, err := DB.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := models.Transaction{}
		rows.Scan(&t.Description, &t.Price, &t.FinalBalance)
		for category, merchants := range models.MerchantMap {
			for _, m := range merchants {
				if strings.Contains(strings.ToLower(strings.Replace(t.Description, " ", "", -1)), m) {
					t.Category = category
				}
			}
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
