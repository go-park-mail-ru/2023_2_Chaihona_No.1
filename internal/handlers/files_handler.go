package handlers

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/gorilla/mux"
)

type FileBody struct {
	Path string `json:"path"`
}

type FileHandler struct {
	Sessions sessions.SessionRepository
	Users users.UserRepository
}

func CreateFileHandler(sessions sessions.SessionRepository, users users.UserRepository) *FileHandler {
	return &FileHandler{
		Sessions: sessions,
		Users: users,
	}
}

func saveFile(fileHeader *multipart.FileHeader, filename string) (string, error) {
	time := time.Now()
	path := filepath.Join(
		// configs.BasePath, 
		"static",
		strconv.Itoa(time.Year()), 
		strconv.Itoa(int(time.Month())),
		strconv.Itoa(time.Day()),
	)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return "", err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	extensions, err := mime.ExtensionsByType(fileHeader.Header.Get("Content-Type"))
	var extension string 
	if err == nil {
		extension = extensions[0]
	}
	path = filepath.Join(path, filename + extension)
	err =  os.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (f *FileHandler) UploadFileStratagy(ctx context.Context, form FileForm) (Result, error){
	fileArray, ok := form.Form.File["avatar"]
	fileBody := FileBody{}
	if ok&&len(fileArray)>0 {
		path, err := saveFile(fileArray[0], "user")
		if err != nil {
			return Result{}, err
		}
		fileBody.Path = path
	}
	return Result{Body: fileBody}, nil
}

func (f *FileHandler) LoadFileStratagy(w http.ResponseWriter, r *http.Request){
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

	// if !auth.CheckAuthorizationByContext(ctx, f.Sessions) {
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
	
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, ErrDataBase.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("avatar", filepath.Base(file.Name()))
	_, err = io.Copy(part, file)
	if err != nil {
		http.Error(w, ErrDataBase.Error(), http.StatusBadRequest)
		return
	}
	writer.Close()

	w.Header().Add("Content-Type", writer.FormDataContentType())
	w.WriteHeader(http.StatusOK)
	//???
	_, err = w.Write(body.Bytes())
	if err != nil {
		log.Println(err)
	}
}