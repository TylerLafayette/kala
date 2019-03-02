package database

import (
	"github.com/TylerLafayette/kala/backend/models"
)

// GetTransactions returns all transactions.
func GetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := "SELECT description, price, final_balance FROM transactions"
	rows, err := DB.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := models.Transaction{}
		rows.Scan(&t.Description, &t.Price, &t.FinalBalance)
		transactions = append(transactions, t)
	}

	return transactions, nil
}
