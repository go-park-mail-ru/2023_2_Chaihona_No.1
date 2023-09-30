package model

const UnpaidReason = "unpaid"
const LowLevelReason = "low_level"

const EveryoneAccess = "for_everyone"
const OneTimePaymentAccess = "one-time_payment"
const SubscribersAccess = "for_subscribers"

type Post struct {
	ID           uint      `json:"id"`
	AuthorID     uint      `json:"-"`
	HasAccess    bool      `json:"has_access"`
	Reason       string    `json:"reason,omitempty"`
	Access       string    `json:"access,string"`
	Payment      float64   `json:"payment,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	MinSubLevel  uint      `json:"min_sub_level,omitempty"`
	CreationDate string    `json:"creation_date,string"`
	Header       string    `json:"header,string"`
	Body         string    `json:"body,omitempty"`
	Likes        uint      `json:"likes"`
	Comments     []Comment `json:"comments,omitempty"`
	Tags         []Tag     `json:"tags"`
}
