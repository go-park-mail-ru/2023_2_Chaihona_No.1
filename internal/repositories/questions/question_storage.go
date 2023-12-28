package questions

import (
	"database/sql"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type QuestionStorage struct {
	db *sql.DB
}

// func SelectQuestionsSQL(id int) squirrel.SelectBuilder {
// 	return squirrel.Select("*").
// 		From("public.question").
// 		Where(squirrel.Eq{"id": id}).
// 		PlaceholderFormat(squirrel.Dollar)
// }

func CreateQuestionStorage(db *sql.DB) *QuestionStorage {
	return &QuestionStorage{
		db: db,
	}
}

func (storage *QuestionStorage) AnswerQuestion(questionId, rating, userId int) error {
	return nil
}

func (storage *QuestionStorage) GetQuestions(userId int) ([]model.Question, error) {
	// rows, err := SelectQuestionsSQL(userId).RunWith(storage.db).Query()

	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// var questions []model.Question
	// err = dbscan.ScanAll(&questions, rows)
	// if err != nil || len(questions) == 0 {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return []model.Question{
		{
			ID: 1,
			Question: "Оцените удобство нашего сервиса",
			QuestionType: 0,
		},
	}, nil
}

func (storage *QuestionStorage) GetStatistic()([]model.Question, error) {
	return []model.Question{
		{
			ID: 1,
			Question: "Оцените удобство нашего сервиса",
			QuestionType: 0,
			Rating: []model.Rating{
				{
					Rating: 1,
					Answers: 3,
				},
				{
					Rating: 2,
					Answers: 7,
				},
				{
					Rating: 3,
					Answers: 19,
				},
				{
					Rating: 4,
					Answers: 59,
				},
				{
					Rating: 5,
					Answers: 12,
				},
			},
		},
		{
			ID: 2,
			Question: "Оцените дизайн нашего сервиса",
			QuestionType: 0,
			Rating: []model.Rating{
				{
					Rating: 1,
					Answers: 3,
				},
				{
					Rating: 2,
					Answers: 17,
				},
				{
					Rating: 3,
					Answers: 39,
				},
				{
					Rating: 4,
					Answers: 59,
				},
				{
					Rating: 5,
					Answers: 12,
				},
			},
		},
	}, nil
}