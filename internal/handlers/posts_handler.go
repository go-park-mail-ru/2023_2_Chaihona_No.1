package handlers

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type PostHandler struct {
	Sessions sessrep.SessionRepository
	Posts    postsrep.PostRepository
}

func CreatePostHandlerViaRepos(session sessrep.SessionRepository, posts postsrep.PostRepository) *PostHandler {
	return &PostHandler{
		session,
		posts,
	}
}

// swagger:route OPTIONS /api/v1/profile/{id}/post Post GetAllUserPostsOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route GET /api/v1/profile/{id}/post Post GetAllUserPosts
// Get user's posts
//
// Parameters:
//   - name: id
//     in: path
//     description: ID of user
//     required: true
//     type: integer
//     format: int
//
// Responses:
//
//	200: result
//	400: result
//	401: result
//	500: result
func (p *PostHandler) GetAllUserPostsStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}

	authorID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	posts, errPost := p.Posts.GetPostsByAuthorId(uint(authorID))

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch err.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				res := Result{Err: errPost.Err.Error()}

				// не понял, зачем нужен закомментированный кусок куда (с присваиванием пустого массива байт)
				// errJSON, err := json.Marshal(res)
				// if err != nil {
				// 	errJSON = []byte{}
				// }
				return res, errPost
			}
			return Result{}, ErrDataBase
		}
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}

	session, _ := p.Sessions.CheckSession(cookie.Value)
	userID := session.UserID
	profile, ok := p.Profiles.GetProfile(uint(userID))
	if !ok {
		return Result{}, ErrNoProfile
	}
	subscriptions := profile.Subscriptions
	isSubscirber := false
	for _, user := range subscriptions {
		if user.ID == uint(authorID) {
			isSubscirber = true
		}
	}

	if !isSubscirber {
		for i := range posts {
			switch posts[i].Access {
			case model.SubscribersAccess:
				posts[i].HasAccess = false
				posts[i].Reason = model.LowLevelReason
				posts[i].Body = ""
			case model.EveryoneAccess:
				posts[i].HasAccess = true
			}
		}
	} else {
		for i := range posts {
			posts[i].HasAccess = true
		}
	}

	return Result{Body: BodyPosts{Posts: posts}}, nil
}
