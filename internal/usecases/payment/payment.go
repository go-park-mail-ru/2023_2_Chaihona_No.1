package payment

import (
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/google/uuid"
)

func Donate(paymentDB model.Payment) (string, string, error) {
	paymnetId, err := uuid.NewRandom()
	if err != nil {
		return "", "", err
	}
	return paymnetId.String(), configs.FakeRedirectURL, nil
}

func requestPaymentStatusAPI() (int, error) {
	return model.PaymentSucceededStatus, nil
}

func CheckPaymentStatusAPI(paymentRepository payments.PaymentRepository, payment model.Payment) {
	status, err := requestPaymentStatusAPI()
	if err != nil {
		log.Println(err)
		return
	}
	if status == model.PaymentWaitingStatus {
		return
	}
	if status == model.PaymentCanceledStatus {
		payment.Status = model.PaymentCanceledStatus
		err = paymentRepository.ChangePayment(payment)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	payment.Status = model.PaymentSucceededStatus
	err = paymentRepository.ChangePayment(payment)
	if err != nil {
		log.Println(err)
	}
}
