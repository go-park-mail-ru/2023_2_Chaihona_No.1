package model

const (
	UnpaidReason         = "unpaid"
	LowLevelReason       = "low_level"
	EveryoneAccess       = "for_everyone"
	OneTimePaymentAccess = "one-time_payment"
	SubscribersAccess    = "for_subscribers"
)

type Post struct {
	ID           uint      `json:"id"`
	AuthorID     uint      `json:"-"`
	HasAccess    bool      `json:"has_access"`
	Reason       string    `json:"reason,omitempty"`
	Access       string    `json:"access"`
	Payment      float64   `json:"payment,omitempty"`
	Currency     string    `json:"currency,omitempty"`
	MinSubLevel  uint      `json:"min_sub_level,omitempty"`
	CreationDate string    `json:"creation_date"`
	Header       string    `json:"header"`
	Body         string    `json:"body,omitempty"`
	Likes        uint      `json:"likes"`
	Comments     []Comment `json:"comments,omitempty"`
	Tags         []Tag     `json:"tags"`
}
