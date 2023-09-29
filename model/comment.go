package model

type Comment struct {
	ID           uint   `json:"id"`
	User         User   `json:"user"`
	Text         string `json:"text,string"`
	CreationDate string `json:"creation_date,string"`
}
