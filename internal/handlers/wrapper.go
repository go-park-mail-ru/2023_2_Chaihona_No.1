package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/gorilla/mux"
)

const (
	maxBytesToRead = 1024 * 2
)

type IValidatable interface {
	IsValide() bool
	IsEmpty() bool
}

type Wrapper[Req IValidatable, Resp any] struct {
	fn func(ctx context.Context, req Req) (Resp, error)
}

// возможно хорошо бы валидатор отдельно сделать, но пока так
func NewWrapper[Req IValidatable, Resp any](fn func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}

func (wrapper *Wrapper[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")

	ctx := auth.AddWriter(r.Context(), w)
	ctx = auth.AddVars(ctx, mux.Vars(r))
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

	rawJSON, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		WriteHttpError(w, ErrEncoding)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawJSON)
}
