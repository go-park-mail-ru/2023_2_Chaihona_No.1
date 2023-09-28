package model

type SubscribeLevel struct {
	ID          uint    `json:"id,uint"`
	Level       uint    `json:"level,uint"`
	Name        string  `json:"name,string"`
	Description string  `json:"description,string"`
	Payment     float64 `json:"payment"`
	Currency    string  `json:"currency,string"`
}
