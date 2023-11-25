package model

type Answer struct {
	ID int `json:"id" db:"id"`
	QuestionId int `json:"question_id" db:"question_id"`
	UserId int `json:"user_id" db:"user_id"`
	Rating int `json:"rating" db:"rating"`
}