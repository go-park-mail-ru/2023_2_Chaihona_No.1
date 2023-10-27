package model

const (
	UnpaidReason         = "unpaid"
	LowLevelReason       = "low_level"
	EveryoneAccess       = "for_everyone"
	OneTimePaymentAccess = "one-time_payment"
	SubscribersAccess    = "for_subscribers"
)

type Post struct {
	ID            uint      `json:"id" db:"id"`
	AuthorID      uint      `json:"-" db:"creator_id"`
	HasAccess     bool      `json:"has_access" db:""`
	Reason        string    `json:"reason,omitempty" db:""`
	Access        string    `json:"access" db:""`
	Payment       float64   `json:"payment,omitempty" db:""`
	Currency      string    `json:"currency,omitempty" db:""`
	MinSubLevel   uint      `json:"min_sub_level,omitempty" db:""`
	MinSubLevelId uint      `json:"-" db:"min_subscription_level_id"`
	CreationDate  string    `json:"creation_date" db:"created_at"`
	Header        string    `json:"header" db:"header"`
	Body          string    `json:"body,omitempty" db:"body"`
	Likes         uint      `json:"likes" db:"likes"`
	Comments      []Comment `json:"comments,omitempty" db:""`
	Tags          []Tag     `json:"tags" db:""`
}
