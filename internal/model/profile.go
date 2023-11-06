package model

type Profile struct {
	User            User             `json:"user"`
	Subscribers     uint             `json:"subscribers,omitempty"`
	SubscribeLevels []SubscribeLevel `json:"subscribe_levels,omitempty"`
	Subscriptions   []User           `json:"subscriptions,omitempty"`
	Donated         string          `json:"donated,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	Goals           []Goal           `json:"goals,omitempty"`
	IsFollowed bool `json:"is_followed"`
}