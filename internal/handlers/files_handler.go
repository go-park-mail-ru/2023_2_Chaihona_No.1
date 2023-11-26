package handlers

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/attaches"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/files"
	"github.com/gorilla/mux"
)

type BodyAttaches struct {
	Attaches []model.Attach `json:"attaches"`
}

type FileBody struct {
	Path string `json:"path"`
}

type FileHandler struct {
	// Sessions sessions.SessionRepository
	// Sessions 
	Attaches attaches.AttachRepository
	Manager *sessions.RedisManager
	Users users.UserRepository
}

func CreateFileHandler(sessionsManager *sessions.RedisManager, users users.UserRepository, attaches attaches.AttachRepository) *FileHandler {
	return &FileHandler{
		// Sessions: sessions,
		Attaches: attaches,
		Manager: sessionsManager,
		Users: users,
	}
}

func (f *FileHandler) UploadFileStratagy(ctx context.Context, form FileForm) (Result, error){
	fileArray, ok := form.Form.File["avatar"]
	fileBody := FileBody{}
	if ok&&len(fileArray)>0 {
		path, err := files.SaveFile(fileArray[0], "user")
		if err != nil {
			return Result{}, err
		}
		fileBody.Path = path
	}
	return Result{Body: fileBody}, nil
}

func (f *FileHandler) LoadFileStratagy(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)

	ctx := auth.AddWriter(r.Context(), w)
	ctx = auth.AddVars(ctx, mux.Vars(r))
	cookie, err := r.Cookie("session_id")
	if err == nil {
		ctx = auth.AddSession(ctx, cookie)
	}
	vars := auth.GetVars(ctx)
	if vars == nil {
		http.Error(w, ErrNoVars.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, ErrBadID.Error(), http.StatusBadRequest)
		return
	}

	// if !auth.CheckAuthorizationManager(ctx, f.Manager) {
	// 	http.Error(w, ErrUnathorized.Error(), http.StatusUnauthorized)
	// 	return
	// }

	user, err := f.Users.GetUser(id)
	if err != nil {
		http.Error(w, ErrDataBase.Error(), http.StatusInternalServerError)
		return
	}
	fileDir, _ := os.Getwd()
	filePath := filepath.Join(fileDir, user.Avatar)

	body, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, ErrDataBase.Error(), http.StatusBadRequest)
		return
	}

	// w.Header().Add("Content-Type", writer.FormDataContentType())
	w.Header().Add("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (f *FileHandler) LoadAttachesStratagy(ctx context.Context, form EmptyForm) (Result, error) {
	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}

	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	attaches, err := f.Attaches.GetPostAttaches(postID)
	if err != nil {
		return Result{}, ErrDataBase
	}

	for i := range attaches {
		fileDir, _ := os.Getwd()
		filePath := filepath.Join(fileDir, attaches[i].FilePath)
	
		body, err := os.ReadFile(filePath)
		if err != nil {
			return Result{}, ErrReadFile
		}
		attaches[i].Data = base64.StdEncoding.EncodeToString(body)
	}
	return Result{Body: BodyAttaches{Attaches: attaches}}, nil
}