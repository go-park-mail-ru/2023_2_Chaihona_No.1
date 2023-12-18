package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	SessionsManager            *sessrep.RedisManager
	PaymentsManager           *paymentsrep.PaymentManager
	Subscriptions      subscriptions.SubscriptionRepository
	SubscriptionLevels subscriptionlevels.SubscribeLevelRepository
}

func CreatePaymentHandlerViaRepos(manager *sessrep.RedisManager,
	payments *paymentsrep.PaymentManager,
	subsciptions subscriptions.SubscriptionRepository,
	subscriptionLevels subscriptionlevels.SubscribeLevelRepository,
) *PaymentHandler {
	return &PaymentHandler{
		SessionsManager:            manager,
		PaymentsManager:           payments,
		Subscriptions:      subsciptions,
		SubscriptionLevels: subscriptionLevels,
	}
}

func (p *PaymentHandler) DonateStratagy(ctx context.Context, form PaymentForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.SessionsManager) {
		return Result{}, ErrUnathorized
	}

	payment := model.Payment{
		DonaterId: form.Body.DonaterId,
		CreatorId: form.Body.CreatorId,
		Currency:  form.Body.Currency,
		Value:     form.Body.Value,
		Type: model.PaymentTypeDonate,
	}
	responseUkassa, err := pay.Donate(payment)
	if err != nil {
		//think
		log.Println(err)
		return Result{}, ErrDataBase
	}

	switch responseUkassa.Status {
	case "pending":
		payment.Status = model.PaymentWaitingStatus
	case "succeeded":
		payment.Status = model.PaymentSucceededStatus
	case "canceled":
		payment.Status = model.PaymentCanceledStatus
	}

	payment.UUID = responseUkassa.Id
	fmt.Println(responseUkassa.Amount.Value)
	splitedValue := strings.Split(responseUkassa.Amount.Value, ".")
	fmt.Println(splitedValue)
	integer, err := strconv.Atoi(splitedValue[0])
	if err != nil {
		//think
		log.Println(err)
		return Result{}, ErrDataBase
	}
	fractional, err := strconv.Atoi(splitedValue[1])
	if err != nil {
		//think
		log.Println(err)
		return Result{}, ErrDataBase
	}
	payment.PaymentInteger = uint(integer)
	payment.PaymentFractional = uint(fractional)
	id, err := p.PaymentsManager.CreateNewPayment(payment)
	if err != nil {
		//think
		log.Println(err)
		return Result{}, ErrDataBase
	}
	payment.Id = id

	c := cron.New()
	_, err = c.AddFunc("* * * * *", func() {
		fmt.Println("joba")
		payment, err := pay.CheckPaymentStatusAPI(p.PaymentsManager, payment)
		if err != nil {
			log.Println(err)
			return
		}
		if payment.Status == model.PaymentSucceededStatus {
			// levels, err := p.SubscriptionLevels.GetUserLevels(payment.CreatorId)
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }
			// for _, level := range levels {
			// 	if level.CostInteger > payment.PaymentInteger {
			// 		continue
			// 	}
			// 	if level.CostInteger == payment.PaymentInteger && level.CostFractional > payment.PaymentFractional {
			// 		continue
			// 	}
			// 	subscription := model.Subscription{
			// 		Subscriber_id:         payment.DonaterId,
			// 		Creator_id:            payment.CreatorId,
			// 		Subscription_level_id: level.ID,
			// 	}
			// 	_, err := p.Subscriptions.AddNewSubscription(subscription)
			// 	if err != nil {
			// 		log.Println(err)
			// 		return
			// 	}
			// 	break
			// }
			c.Stop()
		}
		if payment.Status == model.PaymentCanceledStatus {
			c.Stop()
		}
	})
	c.Start()

	if err != nil {
		//think
		log.Println(err)
		return Result{}, ErrDataBase
	}

	return Result{Body: BodyPayments{RedirectURL: responseUkassa.Confirmation.ConfirmationURL}}, nil
}

func (p *PaymentHandler) GetUsersDonatesStratagy(ctx context.Context, form PaymentForm) (Result, error) {
	// if !auth.CheckAuthorizationManager(ctx, p.Manager) {
	// 	return Result{}, ErrUnathorized
	// }

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	payments, err := p.PaymentsManager.GetPaymentsByUserId(uint(id))
	if err != nil {
		//think
		return Result{}, ErrDataBase
	}
	return Result{Body: BodyPayments{Payments: payments}}, nil
}

func (p *PaymentHandler) GetAuthorDonatesStratagy(ctx context.Context, form PaymentForm) (Result, error) {
	// if !auth.CheckAuthorizationManager(ctx, p.Manager) {
	// 	return Result{}, ErrUnathorized
	// }

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