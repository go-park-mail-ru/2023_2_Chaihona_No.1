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

type Amount struct {
	Value string `json:"value"`
	Currency string `json:"currency"`
} 

type Confirmation struct {
	Type string `json:"type"`
	ReturnURL string `json:"return_url,omitempty"`
	ConfirmationURL string `json:"confirmation_url,omitempty"`
}

type PaymentMethodData struct {
	Type string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Saved bool `json:"saved,omitempty"`
}

type RequestUKassa struct {
	Amount `json:"amount"`
	Capture bool `json:"capture"`
	Confirmation `json:"confirmation"`
	SavePaymentMethod bool `json:"save_payment_method,omitempty"`
	PaymentMethodData `json:"payment_method_data,omitempty"`
}

type Recipient struct {
	AccountId string `json:"account_id"`
	GatewayId string `json:"gateway_id"`
}

type ResponseUKassa struct {
	Id string `json:"id"`
	Status string `json:"status"`
	Paid bool `json:"paid"`
	Amount `json:"amount"`
	Confirmation `json:"confirmation"`
	CreatedAt string `json:"created_at"`
	Description string `json:"description,omitempty"`
	Metadata interface{} `json:"metadata"`
	Recipient `json:"recipient"`
	Refundable bool `json:"refundable"`
	Test bool `json:"test"`
}