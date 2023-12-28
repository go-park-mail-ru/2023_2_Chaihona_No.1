package statistics

import (
	"context"
	"database/sql"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type QeustionsStorage struct {
	db *sql.DB
	UnimplementedStatisticsServer
}

type RedisManager struct {
	statistics StatisticsClient
}

func GetQuestions(userId int) []model.Question {
	return nil
}

func AnswerQuestion(questionId, rating, userId int) {
	return 
}

func GetStatistic() []model.Question {
	return nil
}

func GetQuestionsCtx(ctx context.Context, userId *UserId) (*QuestionsMap, error) {
	return nil, nil
}

func AnswerQuestionCtx(ctx context.Context, answer *Answer) (*Nothing, error) {
	return nil, nil
}

func GetStatisticCtx(ctx context.Context, nothing *Nothing) (*QuestionsMap, error) {
	return nil, nil
}
