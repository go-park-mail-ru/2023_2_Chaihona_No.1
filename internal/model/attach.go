package model

type Attach struct {
	Id       int    `json:"id" db:"id"`
	FilePath string `json:"file_path" db:"file_path"`
}
