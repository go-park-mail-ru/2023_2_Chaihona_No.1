package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	paymentsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	pay "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/payment"
)

type BodyPayments struct {
	RedirectURL string `json:"redirect_url"`
}

type PaymentHandler struct {
	Sessions sessrep.SessionRepository
	Payments paymentsrep.PaymentRepository
}

func CreatePaymentHandlerViaRepos(session sessrep.SessionRepository,
	payments paymentsrep.PaymentRepository,
) *PaymentHandler {
	return &PaymentHandler{
		session,
		payments,
	}
}

func (p *PaymentHandler) Donate(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	payment := &model.Payment{}

	err := decoder.Decode(payment)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	paymentId, redirectURL, err := pay.Donate(*payment)

	payment.Status = model.PaymentWaitingStatus
	payment.UUID = paymentId

	_, err = p.Payments.CreateNewPayment(*payment)
	bodyPayments := BodyPayments{redirectURL}

	result := Result{Body: bodyPayments}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}
