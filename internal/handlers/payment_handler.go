package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/robfig/cron/v3"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	paymentsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	subscriptionlevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscription_levels"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	pay "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/payment"
)

type BodyPayments struct {
	Payments    []model.Payment `json:"payments,omitempty"`
	RedirectURL string          `json:"redirect_url,omitempty"`
}

type PaymentHandler struct {
	// Sessions           sessrep.SessionRepository
	Manager            *sessrep.RedisManager
	Payments           paymentsrep.PaymentRepository
	Subscriptions      subscriptions.SubscriptionRepository
	SubscriptionLevels subscriptionlevels.SubscribeLevelRepository
}

func CreatePaymentHandlerViaRepos(manager *sessrep.RedisManager,
	payments paymentsrep.PaymentRepository,
	subsciptions subscriptions.SubscriptionRepository,
	subscriptionLevels subscriptionlevels.SubscribeLevelRepository,
) *PaymentHandler {
	return &PaymentHandler{
		Manager:            manager,
		Payments:           payments,
		Subscriptions:      subsciptions,
		SubscriptionLevels: subscriptionLevels,
	}
}

func (p *PaymentHandler) DonateStratagy(ctx context.Context, form PaymentForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.Manager) {
		return Result{}, ErrUnathorized
	}
	fmt.Println(form.Body.DonaterId, form.Body.CreatorId)
	payment := model.Payment{
		DonaterId: form.Body.DonaterId,
		CreatorId: form.Body.CreatorId,
		Currency:  form.Body.Currency,
		Value:     form.Body.Value,
	}
	paymentId, redirectURL, err := pay.Donate(payment)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}

	payment.Status = model.PaymentWaitingStatus
	payment.UUID = paymentId
	val, err := strconv.Atoi(payment.Value)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	payment.PaymentInteger = uint(val)
	id, err := p.Payments.CreateNewPayment(payment)
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	payment.Id = id

	c := cron.New()
	_, err = c.AddFunc("* * * * *", func() {
		fmt.Println("joba")
		pay.CheckPaymentStatusAPI(p.Payments, payment)
		if payment.Status != model.PaymentWaitingStatus {
			levels, err := p.SubscriptionLevels.GetUserLevels(payment.CreatorId)
			if err != nil {
				log.Println(err)
				return
			}
			for _, level := range levels {
				if level.CostInteger > payment.PaymentInteger {
					continue
				}
				if level.CostInteger == payment.PaymentInteger && level.CostFractional > payment.PaymentFractional {
					continue
				}
				subscription := model.Subscription{
					Subscriber_id:         payment.DonaterId,
					Creator_id:            payment.CreatorId,
					Subscription_level_id: level.ID,
				}
				_, err := p.Subscriptions.AddNewSubscription(subscription)
				if err != nil {
					log.Println(err)
					return
				}
				break
			}
			c.Stop()
		}
	})
	c.Start()

	if err != nil {
		//think
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyPayments{RedirectURL: redirectURL}}, nil
}

func (p *PaymentHandler) GetUsersDonatesStratagy(ctx context.Context, form PaymentForm) (Result, error) {
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

	payments, err := p.Payments.GetPaymentsByUserId(uint(id))
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	return Result{Body: BodyPayments{Payments: payments}}, nil
}

func (p *PaymentHandler) GetAuthorDonatesStratagy(ctx context.Context, form PaymentForm) (Result, error) {
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

	// payments, err := p.Payments.GetPaymentsByAuthorId(uint(id))
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyPayments{Payments: []model.Payment{}}}, nil
}
