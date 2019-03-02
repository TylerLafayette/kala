package utils

import (
	"strings"

	"github.com/TylerLafayette/kala/backend/models"
)

// GenerateSuggestions generates a list of suggestions based on the user's transaction history.
func GenerateSuggestions(balance, budget float32, transactions []models.Transaction) ([]models.Suggestion, error) {
	suggestions := []models.Suggestion{}
	historyMap := map[string]map[string]interface{}{}
	for _, transaction := range transactions {
		for category, merchants := range models.MerchantMap {
			for _, m := range merchants {
				if strings.Contains(strings.ToLower(strings.Replace(transaction.Description, " ", "", -1)), m) {
					if len(historyMap[category]) > 0 {
						historyMap[category]["count"] = historyMap[category]["count"].(int) + 1
						historyMap[category]["amtSpent"] = historyMap[category]["amtSpent"].(float32) + transaction.Price
					} else {
						historyMap[category] = map[string]interface{}{}
						historyMap[category]["count"] = 1
						historyMap[category]["amtSpent"] = transaction.Price
					}
				}
			}
		}
	}

	for k, v := range historyMap {
		if v["amtSpent"].(float32) > 15 || v["count"].(int) > 2 {
			s := models.Suggestions[k]
			s.AmountSpentThisMonth = v["amtSpent"].(float32)
			suggestions = append(suggestions, s)
		}
	}

	return suggestions, nil
}
