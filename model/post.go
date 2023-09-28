package model

const UnpaidReason = "unpaid"
const LowLevelReason = "low_level"

const EveryoneAccess = "for_everyone"
const OneTimePaymentAccess = "one-time_payment"
const SubscribersAccess = "for_subscribers"

type Post struct {
	ID           uint      `json:"id,uint"`
	HasAccess    bool      `json:"has_access,bool"`
	Reason       string    `json:"reason,omitempty"`
	Access       string    `json:"access,string"`
	Payment      float64   `json:"payment"`
	MinSubLevel  uint      `json:"min_sub_level"`
	CreationDate string    `json:"creation_date"`
	Header       string    `json:"header"`
	Body         string    `json:"body,omitempty"`
	Likes        uint      `json:"likes"`
	Comments     []Comment `json:"comments,omitempty"`
	Tags         []Tag     `json:"tags"`
}
