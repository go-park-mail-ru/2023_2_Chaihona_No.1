package model

type Attach struct {
	Id       int    `json:"id" db:"id"`
	PostId int `json:"post_id" db:"post_id"`
	FilePath string `json:"file_path" db:"file_path"`
	Name string `json:"name" db:"name"`
	Data string `json:"data" db:""`
}