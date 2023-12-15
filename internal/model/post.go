package model

import "database/sql"

const (
	UnpaidReason         = "unpaid"
	LowLevelReason       = "low_level"
	EveryoneAccess       = "for_everyone"
	OneTimePaymentAccess = "one-time_payment"
	SubscribersAccess    = "for_subscribers"
)

type Post struct {
	ID            uint      `json:"id" db:"id"`
	AuthorID      uint      `json:"creator_id" db:"creator_id"`
	HasAccess     bool      `json:"has_access" db:"has_access"`
	Reason        string    `json:"reason,omitempty" db:""`
	// Access        string    `json:"access" db:""`
	Payment       float64   `json:"payment,omitempty" db:""`
	Currency      string    `json:"currency,omitempty" db:""`
	MinSubLevel   uint      `json:"min_sub_level" db:"min_sub_level"`
	MinSubLevelId uint      `json:"-" db:"min_subscription_level_id"`
	CreationDate  string    `json:"created_at" db:""`
	CreationDateSQL sql.NullTime `json:"-" db:"created_at"`
	UpdatedAt     string    `json:"updated_at" db:"updated_at"`
	Header        string    `json:"header" db:"header"`
	Body          string    `json:"body,omitempty" db:"body"`
	Likes         uint      `json:"likes" db:"likes"`
	Comments      []Comment `json:"comments,omitempty" db:"comments"`
	Tags          []Tag     `json:"tags" db:""`
	Attaches      string    `json:"attaches" db:"attaches"`
	IsLiked bool `json:"is_liked" db:"is_liked"`
}