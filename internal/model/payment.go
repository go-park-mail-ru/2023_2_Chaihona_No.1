package model

const (
	PaymentWaitingStatus   = 0
	PaymentCanceledStatus  = 1
	PaymentSucceededStatus = 2
)

type Payment struct {
	Id int `json:"-" db:"id"`
	UUID              string `json:"-" db:"uuid"`
	PaymentInteger    uint   `json:"-" db:"payment_integer"`
	PaymentFractional uint   `json:"-" db:"payment_fractional"`
	Status            uint   `json:"-" db:"status"`
	DonaterId         uint   `json:"donater_id" db:"donater_id"`
	CreatorId         uint   `json:"creator_id" db:"creator_id"`
	Currency          string `json:"currency,omitempty" db:"currency"`
	Value             string `json:"value,omitempty" db:""`
}