package payments

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type PaymentRepository interface {
	CreateNewPayment(payment model.Payment) (int, error)
	DeletePayment(id uint) error
	GetPayment(uuid string) (model.Payment, error)
	GetPaymentsByAuthorId(authorID uint) (model.Payment, error)
	GetPaymentsByUserId(userId uint) ([]model.Payment, error)
	ChangePayment(payment model.Payment) error
}