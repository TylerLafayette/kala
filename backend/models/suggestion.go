package models

// Suggestion represents a suggestion based on a transaction history.
type Suggestion struct {
	Emoji                string      `json:"emoji"`
	Header               string      `json:"header"`
	Description          string      `json:"description"`
	AmountSpentThisMonth float32     `json:"amountSpentThisMonth"`
	CallToActionType     string      `json:"callToActionType"`
	Data                 interface{} `json:"data"`
}
