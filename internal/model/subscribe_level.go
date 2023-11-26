package model

type SubscribeLevel struct {
	ID             uint    `json:"id" db:"id"`
	Level          uint    `json:"level" db:"level"`
	Name           string  `json:"name" db:"name"`
	Description    string  `json:"description" db:"description"`
	// Payment        float64 `json:"payment" db:""`
	Currency       string  `json:"currency" db:"currency"`
	CostInteger    uint    `json:"cost_integer" db:"cost_integer"`
	CostFractional uint    `json:"cost_fractional" db:"cost_fractional"`
	CreationDate   string  `json:"-" db:"created_at"`
	LastUpdate     string  `json:"-" db:"updated_at"`
	CreatorID      uint    `json:"-" db:"creator_id"`
}