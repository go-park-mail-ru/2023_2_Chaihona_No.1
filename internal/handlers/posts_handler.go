package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/attaches"
	likesrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/files"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type BodyLike struct {
	PostId int `json:"post_id"`
}

type PostHandler struct {
	// Sessions sessrep.SessionRepository
	Manager *sessrep.RedisManager
	Posts   postsrep.PostRepository
	Likes   likesrep.LikeRepository
	Attaches attaches.AttachRepository
}

func CreatePostHandlerViaRepos(manager *sessrep.RedisManager, posts postsrep.PostRepository,
	likes likesrep.LikeRepository, attaches attaches.AttachRepository) *PostHandler {
	return &PostHandler{
		manager,
		posts,
		likes,
		attaches,
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
	// if !auth.CheckAuthorizationManager(ctx, p.Manager) {
	// 	return Result{}, ErrUnathorized
	// }

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}

	authorID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	queryVars := auth.GetQueryVars(ctx)
	if queryVars == nil {
		return Result{}, ErrNoVars
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, _ := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	var posts []model.Post
	var errPost error
	if queryVars["is_owner"] == "true" {
		//check authorization
		posts, errPost = p.Posts.GetOwnPostsByAuthorId(uint(authorID), uint(authorID))
	} else {
		if queryVars["is_followed"] == "true" {
			//check authorization
			posts, errPost = p.Posts.GetPostsByAuthorIdForFollower(uint(authorID), uint(session.UserID))
		} else {
			posts, errPost = p.Posts.GetPostsByAuthorIdForStranger(uint(authorID), uint(session.UserID))
		}
	}
	for i := range posts {
		posts[i].CreationDate = posts[i].CreationDateSQL.Time.Format("2006-01-02 15:04")
	}
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
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
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

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}
	post := model.Post{
		ID:            uint(postId),
		MinSubLevelId: form.Body.MinSubLevelId,
		Header:        form.Body.Header,
		Body:          form.Body.Body,
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	post.AuthorID = uint(session.UserID)
	err = p.Posts.ChangePost(post)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}

	for _, deleted_path := range form.Body.Pinned.Deleted {
		err = p.Attaches.DeleteAttach(deleted_path)
		if err != nil {
			return Result{}, ErrDataBase
		}

		err = files.DeleteFile(deleted_path);
		if err != nil {
			log.Println(err)
			return Result{}, ErrDeleteFile
		}
	}

	for i, attach := range form.Body.Pinned.Files {
		countedAttaches, err := p.Attaches.CountAttaches(postId)
		if err != nil {
			return Result{}, ErrDataBase
		}

		path, err := files.SaveFileBase64(attach.Data, fmt.Sprintf("attach%d_post%d%s", countedAttaches + i, postId, attach.Name[strings.LastIndexByte(attach.Name, '.'):]))
		if err != nil {
			log.Println(err)
			return Result{}, ErrSaveFile
		}
		_, err = p.Attaches.PinAttach(model.Attach{
			PostId: postId,
			FilePath: path,
			Name: attach.Name,
		})
		if err != nil {
			return Result{}, ErrDataBase
		}
	}
	return Result{}, nil
}

func (p *PostHandler) CreateNewPostStrategy(ctx context.Context, form PostForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	post := model.Post{
		AuthorID:      uint(session.UserID),
		MinSubLevelId: form.Body.MinSubLevelId,
		Header:        form.Body.Header,
		Body:          form.Body.Body,
	}
	postId, err := p.Posts.CreateNewPost(post)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	
	for i, attach := range form.Body.Attaches {
		//check extension
		path, err := files.SaveFileBase64(attach.Data, fmt.Sprintf("attach%d_post%d%s", i, postId, attach.Name[strings.LastIndexByte(attach.Name, '.'):]))
		if err != nil {
			log.Println(err)
			return Result{}, ErrSaveFile
		}
		_, err = p.Attaches.PinAttach(model.Attach{
			PostId: postId,
			FilePath: path,
			Name: attach.Name,
		})
		if err != nil {
			return Result{}, ErrDataBase
		}
	}

	bodyResponse := map[string]interface{}{
		"id": postId,
	}
	return Result{Body: bodyResponse}, nil
}

func (p *PostHandler) DeletePostStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
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

//Добавить обработку для ананоимного ползователя
func (p *PostHandler) GetFeedStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	// if !auth.CheckAuthorizationManager(ctx, p.Manager) {
	// 	return Result{}, ErrUnathorized
	// }

	cookie := auth.GetSession(ctx)
	if cookie == nil {
		return Result{}, ErrNoCookie
	}
	session, ok := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	posts, err := p.Posts.GetUsersFeed(uint(session.UserID))
	if err != nil {
		fmt.Println(err)
		return Result{}, ErrDataBase
	}
	for i := range posts {
		posts[i].CreationDate = posts[i].CreationDateSQL.Time.Format("2006-01-02 15:04")
	}
	return Result{Body: BodyPosts{Posts: posts}}, nil
}

func (p *PostHandler) LikePostStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
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
	session, ok := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
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
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
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
	session, ok := p.Manager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err = p.Likes.DeleteLike(int(session.UserID), id)
	if err != nil {
		return Result{}, ErrDataBase
	}

	return Result{}, nil
}