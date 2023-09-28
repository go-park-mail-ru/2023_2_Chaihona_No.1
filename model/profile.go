package model

type Profile struct {
	ID              uint             `json:"id,int"`
	User            User             `json:"user,User"`
	Description     string           `json:"description,omitempty"`
	Subscribers     uint             `json:"subscribers,omitempty"`
	Goals           []Goal           `json:"goals,omitempty"`
	SubscribeLevels []SubscribeLevel `json:"subscribe_levels,omitempty"`
	Subsribtions    []User           `json:"subsribtions,omitempty"`
	Donated         float64          `json:"donated,omitempty"`
}
