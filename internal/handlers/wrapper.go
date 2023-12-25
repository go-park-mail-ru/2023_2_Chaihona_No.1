package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"mime"
	"mime/multipart"
	"net/http"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	jwriter "github.com/mailru/easyjson/jwriter"
)

const (
	maxBytesToRead = 1024 * 1024 * 1024
)

type IValidatable interface {
	IsValide() bool
	IsEmpty() bool
}

type IMarshable interface {
	MarshalJSON() ([]byte, error) 
	MarshalEasyJSON(w *jwriter.Writer)
}

type Wrapper[Req IValidatable, Resp IMarshable] struct {
	fn func(ctx context.Context, req Req) (Resp, error)
}

// возможно хорошо бы валидатор отдельно сделать, но пока так
func NewWrapper[Req IValidatable, Resp IMarshable](fn func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}


func (wrapper *Wrapper[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	
	ctx := auth.AddWriter(r.Context(), w)
	ctx = auth.AddVars(ctx, mux.Vars(r))

	isOwner := r.URL.Query().Get("is_owner")
	isFollowed := r.URL.Query().Get("is_followed")
	ctx = auth.AddQueryVars(ctx, map[string]string{
		"is_owner": isOwner,
		"is_followed": isFollowed,
	})

	cookie, err := r.Cookie("session_id")
	if err == nil {
		ctx = auth.AddSession(ctx, cookie)
	}
	
	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)
	

	var request Req
	if !request.IsEmpty() {
		err = json.NewDecoder(body).Decode(&request)
		if err != nil {
			WriteHttpError(w, ErrDecoding)
			log.Println(err)
			return
		}

		if !request.IsValide() {
			WriteHttpError(w, ErrValidation)
			return
		}
	}

	response, err := wrapper.fn(ctx, request)
	if err != nil {
		log.Printf("%s: error: %v\n", r.URL, err)
		WriteHttpError(w, err)
		return
	}

	rawJSON, err := easyjson.Marshal(response)
	if err != nil {
		log.Println(err)
		WriteHttpError(w, ErrEncoding)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rawJSON)
	if err != nil {
		log.Println(err)
	}
}

type FileWrapper[Resp IMarshable] struct {
	fn func(ctx context.Context, req FileForm) (Resp, error)
}

func NewFileWrapper[Resp IMarshable](fn func(ctx context.Context, req FileForm) (Resp, error)) *FileWrapper[Resp] {
	return &FileWrapper[Resp]{
		fn: fn,
	}
}

func (wrapper *FileWrapper[Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")

	ctx := auth.AddWriter(r.Context(), w)
	ctx = auth.AddVars(ctx, mux.Vars(r))
	cookie, err := r.Cookie("session_id")
	if err == nil {
		ctx = auth.AddSession(ctx, cookie)
	}
	// //????
	// err = r.ParseMultipartForm(math.MaxInt64)
	// if err != nil {
	// 	WriteHttpError(w, err)
	// }
	contentTypeHeader := r.Header.Get("Content-Type")
	mediaType, params, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		WriteHttpError(w, err)
	}
	fmt.Println(mediaType, params)
	boundary, ok := params["boundary"]
	if !ok {
		WriteHttpError(w, fmt.Errorf("%s", "no boundary"))
	}
	body := http.MaxBytesReader(w, r.Body, r.ContentLength)
	form, err := multipart.NewReader(body, boundary).ReadForm(math.MaxInt64)
	if err != nil {
		WriteHttpError(w, err)
	}
	
	response, err := wrapper.fn(ctx, FileForm{Form: *form})
	if err != nil {
		log.Printf("%s: error: %v\n", r.URL, err)
		WriteHttpError(w, err)
		return
	}

	rawJSON, err := easyjson.Marshal(response)
	if err != nil {
		log.Println(err)
		WriteHttpError(w, ErrEncoding)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rawJSON)
	if err != nil {
		log.Println(err)
	}
}