package model

const (
	SubcribersGoalType = "subscribers"
	MoneyGoalType      = "money"
	RubCurrency        = "rub"
)

type Goal struct {
	ID          uint    `json:"id"`
	GoalType    string  `json:"goal_type"`
	Currency    string  `json:"currency,omitempty"`
	Current     float64 `json:"current"`
	GoalValue   float64 `json:"goal_value"`
	Description string  `json:"description"`
}
