package model

const (
	FiveStarsType = 0
)

type Rating struct {
	Rating int `json:"rating"`
	Answers int `json:"answers"`
}

type Question struct {
	ID uint `json:"id" db:"id"`
	Question string `json:"question" db:"question"`
	QuestionType uint `json:"question_type" db:"question_type"`
	Rating []Rating `json:"rating"`
}