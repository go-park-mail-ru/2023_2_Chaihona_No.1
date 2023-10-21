package model

type SubscribeLevel struct {
	ID          uint    `json:"id"`
	Level       uint    `json:"level"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Payment     float64 `json:"payment"`
	Currency    string  `json:"currency"`
}
