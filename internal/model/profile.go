package model

type Profile struct {
	User            User             `json:"user"`
	Description     string           `json:"description,omitempty"`
	Subscribers     uint             `json:"subscribers,omitempty"`
	Goals           []Goal           `json:"goals,omitempty"`
	SubscribeLevels []SubscribeLevel `json:"subscribe_levels,omitempty"`
	Subscriptions   []User           `json:"subscriptions,omitempty"`
	Donated         float64          `json:"donated,omitempty"`
	Currency        string           `json:"currency,omitempty"`
}