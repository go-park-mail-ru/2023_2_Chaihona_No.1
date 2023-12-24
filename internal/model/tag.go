package model

type Tag struct {
	ID   uint   `json:"id" db:"id"`
	PostId uint `json:"post_id" db:"post_id"`
	Name string `json:"name" db:"name"`
}