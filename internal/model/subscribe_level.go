package model

type SubscribeLevel struct {
	ID              uint    `json:"id" db:"id"`
	Level           uint    `json:"level" db:"level"`
	Name            string  `json:"name" db:"name"`
	Description     string  `json:"description" db:"description"`
	Payment         float64 `json:"payment" db:""`
	Currency        string  `json:"currency" db:"currency"`
	Cost_integer    uint    `json:"-" db:"cost_integer"`
	Cost_fractional uint    `json:"-" db:"cost_fractional"`
	CreationDate    string  `json:"-" db:"creation_date"`
	LastUpdate      string  `json:"-" db:"last_update"`
	CreatorID       uint    `json:"-" db:"creator_id"`
}
