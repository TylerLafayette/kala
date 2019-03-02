package models

// Overview provides a struct to be used by the API overview endpoint.
type Overview struct {
	Username       string        `json:"username"`
	CurrentBalance float32       `json:"currentBalance"`
	Budget         float32       `json:"budegt"`
	Transactions   []Transaction `json:"transactions"`
	Suggestions    []Suggestion  `json:"suggestions"`
}
