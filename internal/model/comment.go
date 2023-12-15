package model

type Comment struct {
	ID           uint   `json:"id" db:"id"`
	UserId         int   `json:"user_id" db:"user_id"`
	PostId	int `json:"post_id" db:"post_id"`
	Text         string `json:"text" db:"text"`
	CreationDate string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}