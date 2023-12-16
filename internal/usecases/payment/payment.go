package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/google/uuid"
)

func Donate(paymentDB model.Payment) (model.ResponseUKassa, error) {
	paymnetId, err := uuid.NewRandom()
	if err != nil {
		return  model.ResponseUKassa{} ,err
	}
	paymentDB.UUID = paymnetId.String();
	requestUkassa := model.RequestUKassa{
		Amount: model.Amount{
			Value: paymentDB.Value,
			Currency: paymentDB.Currency,
		},
		Capture: true,
		Confirmation: model.Confirmation{
			Type: "redirect",
			ReturnURL: "https://my-kopilka.ru",
		},
	}
	requestJson, err := json.Marshal(requestUkassa)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 
	reader := bytes.NewReader(requestJson)
	req, err := http.NewRequest("POST", configs.PaymentAPI + "payments", reader)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", paymentDB.UUID)
	fileDir, _ := os.Getwd()
	filePath := filepath.Join(fileDir, "API_key")
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 
	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(configs.ShopId + ":" + string(file))))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	}
	defer resp.Body.Close()

	var responseUKassa model.ResponseUKassa
	err = json.NewDecoder(resp.Body).Decode(&responseUKassa)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	}
	return responseUKassa, nil
}

func requestPaymentStatusAPI(payment model.Payment) (int, error) {
	req, err := http.NewRequest("GET", configs.PaymentAPI + "payments/" + payment.UUID, nil)
	if err != nil {
		log.Println(err)
		return 0, err
	} 
	// req.Header.Set("Idempotence-Key", payment.UUID)
	fileDir, _ := os.Getwd()
	filePath := filepath.Join(fileDir, "API_key")
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return 0, err
	} 
	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(configs.ShopId + ":" + string(file))))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer resp.Body.Close()

	var responseUKassa model.ResponseUKassa
	err = json.NewDecoder(resp.Body).Decode(&responseUKassa)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	switch responseUKassa.Status {
	case "pending":
		return model.PaymentWaitingStatus, nil
	case "succeeded":
		return model.PaymentSucceededStatus, nil
	case "canceled":
		return model.PaymentCanceledStatus, nil
	case "waiting_for_capture":
		return model.PaymentSucceededStatus, nil
	}

	return 0, errors.New("unexpected behaviour")
}

func CheckPaymentStatusAPI(paymentRepository payments.PaymentRepository, payment model.Payment) (model.Payment, error) {
	status, err := requestPaymentStatusAPI(payment)
	if err != nil {
		log.Println(err)
		return model.Payment{}, err
	}
	payment.Status = uint(status);
	switch status {
		case model.PaymentWaitingStatus:
			return payment, nil
		case model.PaymentCanceledStatus:
			err = paymentRepository.ChangePayment(payment)
			if err != nil {
				log.Println(err)
				return model.Payment{}, err
			}
			return payment, nil
		default:
			err = paymentRepository.ChangePayment(payment)
			if err != nil {
				log.Println(err)
				return model.Payment{}, err
			}
			return payment, nil
	}
}