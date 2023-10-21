package model

type Comment struct {
	ID           uint   `json:"id"`
	User         User   `json:"user"`
	Text         string `json:"text"`
	CreationDate string `json:"creation_date"`
}
