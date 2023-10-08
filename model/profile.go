package model

type Profile struct {
	ID              uint             `json:"id"`
	User            User             `json:"user"`
	Description     string           `json:"description,omitempty"`
	Subscribers     uint             `json:"subscribers,omitempty"`
	Goals           []Goal           `json:"goals,omitempty"`
	SubscribeLevels []SubscribeLevel `json:"subscribe_levels,omitempty"`
	Subscribtions    []User           `json:"subsribtions,omitempty"`
	Donated         float64          `json:"donated,omitempty"`
}
