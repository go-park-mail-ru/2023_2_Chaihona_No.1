package model

const SubcribersGoalType = "subscribers"
const MoneyGoalType = "money"

const RubCurrency = "rub"

type Goal struct {
	ID          uint    `json:"id,uint"`
	GoalType    string  `json:"goal_type,string"`
	Currency    string  `json:"currency,string"`
	Current     float64 `json:"current"`
	GoalValue   float64 `json:"goal_value"`
	Description string  `json:"description,string"`
}
