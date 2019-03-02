package models

// Suggestions contains all the suggestions that can be made.
var Suggestions = map[string]Suggestion{
	"coffee": Suggestion{
		Emoji:            "☕️",
		Header:           "Cut down on the coffee.",
		Description:      "Maybe try buying cheaper instant coffee that delivers the same caffeine boost.",
		CallToActionType: "SAVING_SCHEDULE",
		Data: map[string]interface{}{
			"amtMonthly": 5.0,
		},
	},
	"groceries": Suggestion{
		Emoji:            "🛒",
		Header:           "Shop at cheaper locations.",
		Description:      "Your grocery spending is very high, consider going to a cheaper store.",
		CallToActionType: "SAVING_SCHEDULE",
		Data: map[string]interface{}{
			"amtMonthly": 5.0,
		},
	},
	"restaurants": Suggestion{
		Emoji:            "🍔",
		Header:           "Try cooking at home more.",
		Description:      "You seem to be going to restaurants often. Why not try making your own food at home?",
		CallToActionType: "LINK_BUTTON",
		Data: map[string]interface{}{
			"url":  "https://www.allrecipes.com/recipes/",
			"text": "FIND RECIPES",
		},
	},
}

// MerchantMap stores a list of all merchants, based on category.
var MerchantMap = map[string][]string{
	"coffee": []string{
		"coffeebar",
		"starbucks",
		"sbucks",
		"philzcoffee",
	},
	"groceries": []string{
		"walgreens",
		"cvs",
		"safeway",
	},
	"restaurants": []string{
		"britishbankerscl",
		"mcdonald",
		"subway",
		"olivegarden",
	},
}
