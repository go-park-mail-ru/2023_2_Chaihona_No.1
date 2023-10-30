package payment

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/google/uuid"
)

func Donate(paymentDB model.Payment) (string, string, error) {
	paymnetId, err := uuid.NewRandom()
	if err != nil {
		return "", "", err
	}
	return paymnetId.String(), configs.FakeRedirectURL, nil
}
