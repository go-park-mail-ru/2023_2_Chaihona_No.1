package statistics

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type StatisticRepository interface {
	GetQuestions(userId int) []model.Question
	AnswerQuestion(questionId, rating, userId int)
	GetStatistic() []model.Question
}
