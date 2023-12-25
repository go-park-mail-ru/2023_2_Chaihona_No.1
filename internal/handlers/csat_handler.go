package handlers

import (
	"context"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	questionsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/questions"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
)

//easyjson:json
type BodyQuestions struct{
	Questions []model.Question `json:"questions"`
}

type CSATHandler struct {
	Manager *sessrep.RedisManager
	Questions questionsrep.QuestionRepository
	// Answers answersrep.AnswerRepository
}

func CreateCSATHandler(manager *sessrep.RedisManager, questions questionsrep.QuestionRepository,
	// answers answersrep.AnswerRepository
	) *CSATHandler {
		return &CSATHandler{
			manager,
			questions,
			// answers,
		}
}

func (h *CSATHandler) GetQuestions(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, h.Manager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := h.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	questions, err := h.Questions.GetQuestions(int(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyQuestions{Questions: questions}}, nil
}

func (h *CSATHandler) Rate(ctx context.Context, form RatingForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, h.Manager) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}

	questionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := h.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	err = h.Questions.AnswerQuestion(questionID, form.Body.Rating ,int(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{}, nil
}

func (h *CSATHandler) GetStatistic(ctx context.Context, form EmptyForm) (Result, error) {
	questions, err := h.Questions.GetStatistic()
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyQuestions{Questions: questions}}, nil
}