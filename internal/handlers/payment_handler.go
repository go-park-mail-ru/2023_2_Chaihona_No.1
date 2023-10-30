package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	paymentsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	pay "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/payment"
)

type BodyPayments struct {
	Payments    []model.Payment `json:"payments,omitempty"`
	RedirectURL string          `json:"redirect_url,omitempty"`
}

type PaymentHandler struct {
	Sessions       sessrep.SessionRepository
	Payments       paymentsrep.PaymentRepository
	Susbscriptions subscriptions.SubscriptionRepository
}

func CreatePaymentHandlerViaRepos(session sessrep.SessionRepository,
	payments paymentsrep.PaymentRepository,
	subsciptions subscriptions.SubscriptionRepository,
) *PaymentHandler {
	return &PaymentHandler{
		Sessions:       session,
		Payments:       payments,
		Susbscriptions: subsciptions,
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
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	payment.Status = model.PaymentWaitingStatus
	payment.UUID = paymentId

	_, err = p.Payments.CreateNewPayment(*payment)
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	c := cron.New()
	c.AddFunc("* * * * *", func() {
		pay.CheckPaymentStatusAPI(p.Payments, *payment)
		if payment.Status != model.PaymentWaitingStatus {
			c.Stop()
		}
	})

	bodyPayments := BodyPayments{RedirectURL: redirectURL}

	result := Result{Body: bodyPayments}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}

func (p *PaymentHandler) GetUsersDonates(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	payments, err := p.Payments.GetPaymentsByUserId(uint(userId))
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	result := Result{Body: BodyPayments{Payments: payments}}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}

func (p *PaymentHandler) GetAuthorDonates(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	authorId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	payments, err := p.Payments.GetPaymentsByAuthorId(uint(authorId))
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	result := Result{Body: BodyPayments{Payments: payments}}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}
