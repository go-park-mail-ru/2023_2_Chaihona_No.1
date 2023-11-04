package handlers

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	likesrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type BodyLike struct {
	PostId int `json:"post_id"`
}

type PostHandler struct {
	Sessions sessrep.SessionRepository
	Posts    postsrep.PostRepository
	Likes    likesrep.LikeRepository
}

func CreatePostHandlerViaRepos(session sessrep.SessionRepository, posts postsrep.PostRepository,
	likes likesrep.LikeRepository) *PostHandler {
	return &PostHandler{
		session,
		posts,
		likes,
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

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	posts, errPost := p.Posts.GetPostsByAuthorId(uint(authorID), uint(session.UserID))

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch err.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				return Result{}, errPost
			}
			return Result{}, ErrDataBase
		}
		//подумать
		return Result{}, errPost
	}

	return Result{Body: BodyPosts{Posts: posts}}, nil
}

func (p *PostHandler) ChangePostStrategy(ctx context.Context, form PostForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	post := model.Post{
		ID: form.Body.ID,
		MinSubLevelId: form.Body.MinSubLevelId,
		Header: form.Body.Header,
		Body: form.Body.Body,
		Likes: form.Body.Likes,
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	post.AuthorID = uint(session.UserID)
	err = p.Posts.ChangePost(post)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *PostHandler) CreateNewPostStrategy(ctx context.Context, form PostForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	post := model.Post{
		ID: form.Body.ID,
		MinSubLevelId: form.Body.MinSubLevelId,
		Header: form.Body.Header,
		Body: form.Body.Body,
		Likes: form.Body.Likes,
	}

	postId, err := p.Posts.CreateNewPost(post)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}

	bodyResponse := map[string]interface{}{
		"id": postId,
	}
	return Result{Body: bodyResponse}, nil
}

func (p *PostHandler) DeletePostStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	err = p.Posts.DeletePost(uint(id))
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *PostHandler) GetFeedStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	
	posts, err := p.Posts.GetUsersFeed(uint(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyPosts{Posts: posts}}, nil
}

func (p *PostHandler) LikePostStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err = p.Likes.CreateNewLike(int(session.UserID), id)
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{}, nil
}

func (p *PostHandler) UnlikePostStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Sessions) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err = p.Likes.DeleteLike(int(session.UserID), id)
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{}, nil
}