package questions

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type QuestionRepository interface {
	GetQuestions(userId int) ([]model.Question, error)
	AnswerQuestion(questionId, rating, userId int) error
	GetStatistic() ([]model.Question, error)
}