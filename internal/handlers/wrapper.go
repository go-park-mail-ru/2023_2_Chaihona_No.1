package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

const (
	maxBytesToRead = 1024 * 2
)

var (
	validationErr = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"user_validation"}`,
	}

	decodingErr = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"wrong_json"}`,
	}

	encodingErr = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"encoding_json"}`,
	}
)

type IValidatable interface {
	IsValide() bool
}

type Wrapper[Req IValidatable, Resp any] struct {
	fn func(ctx context.Context, req Req) (Resp, error)
}

func NewWrapper[Req IValidatable, Resp any](fn func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		fn: fn,
	}
}

func (wrapper *Wrapper[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")

	ctx := r.Context()

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	var request Req
	err := json.NewDecoder(body).Decode(&request)
	if err != nil {
		WriteHttpError(w, decodingErr)
		return
	}

	if !request.IsValide() {
		WriteHttpError(w, validationErr)
		return
	}

	response, err := wrapper.fn(ctx, request)
	if err != nil {
		log.Printf("%s: error: %v\n", r.URL, err)
		WriteHttpError(w, err)
		return
	}

	rawJSON, err := json.Marshal(response)
	if err != nil {
		WriteHttpError(w, encodingErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawJSON)
}
