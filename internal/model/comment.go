package model

type Comment struct {
	ID           uint   `json:"id"`
	UserId         int   `json:"user_id"`
	PostId	int `json:"post_id"`
	Text         string `json:"text"`
	CreationDate string `json:"created_at"`
}