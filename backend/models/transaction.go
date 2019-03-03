package models

// Transaction represents a single transaction.
type Transaction struct {
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Price        float32 `json:"price"`
	FinalBalance float32 `json:"finalBalance"`
}
