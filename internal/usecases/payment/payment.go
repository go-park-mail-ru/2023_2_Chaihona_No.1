package payment

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

func createAutoPaymentUKassa(payment model.Payment) (model.Payment, model.SubRequestUKassa, error) {
	paymnetId, err := uuid.NewRandom()
	if err != nil {
		return  model.Payment{}, model.SubRequestUKassa{}, err
	}
	payment.UUID = paymnetId.String();
	requestUkassa := model.SubRequestUKassa{
		Amount: model.Amount{
			Value: payment.Value,
			Currency: payment.Currency,
		},
		Capture: true,
	}
	return payment, requestUkassa, nil
}

func createPaymentUKassa(payment model.Payment, save bool) (model.Payment, model.RequestUKassa, error) {
	paymnetId, err := uuid.NewRandom()
	if err != nil {
		return  model.Payment{}, model.RequestUKassa{}, err
	}
	payment.UUID = paymnetId.String();
	requestUkassa := model.RequestUKassa{
		Amount: model.Amount{
			Value: payment.Value,
			Currency: payment.Currency,
		},
		Capture: true,
		Confirmation: model.Confirmation{
			Type: "redirect",
			ReturnURL: "https://my-kopilka.ru",
		},
	}
	if save {
		requestUkassa.SavePaymentMethod = true
		requestUkassa.PaymentMethodData = model.PaymentMethodData{
			Type: "bank_card",
		}
	}
	return payment, requestUkassa, nil
}

func makeRequestUKassa(reader io.Reader, method string, UUID string) (model.ResponseUKassa, error) {
	var url string
	switch method {
	case "GET":
		url = configs.PaymentAPI + "payments/" + UUID
		reader = nil
	case "POST":
		url = configs.PaymentAPI + "payments"
	}
	
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotence-Key", UUID)
	}

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

func Donate(paymentDB model.Payment) (model.ResponseUKassa, error) {
	// paymnetId, err := uuid.NewRandom()
	// if err != nil {
	// 	return  model.ResponseUKassa{} ,err
	// }
	// paymentDB.UUID = paymnetId.String();
	// requestUkassa := model.RequestUKassa{
	// 	Amount: model.Amount{
	// 		Value: paymentDB.Value,
	// 		Currency: paymentDB.Currency,
	// 	},
	// 	Capture: true,
	// 	Confirmation: model.Confirmation{
	// 		Type: "redirect",
	// 		ReturnURL: "https://my-kopilka.ru",
	// 	},
	// }
	paymentDB, requestUkassa, err := createPaymentUKassa(paymentDB, false)
	if err != nil {
		return  model.ResponseUKassa{} ,err
	}
	requestJson, err := json.Marshal(requestUkassa)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 
	reader := bytes.NewReader(requestJson)
	responseUKassa, err := makeRequestUKassa(reader, "POST", paymentDB.UUID)
		if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	}
	// req, err := http.NewRequest("POST", configs.PaymentAPI + "payments", reader)
	// if err != nil {
	// 	log.Println(err)
	// 	return model.ResponseUKassa{}, err
	// } 
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Idempotence-Key", paymentDB.UUID)
	// fileDir, _ := os.Getwd()
	// filePath := filepath.Join(fileDir, "API_key")
	// file, err := os.ReadFile(filePath)
	// if err != nil {
	// 	log.Println(err)
	// 	return model.ResponseUKassa{}, err
	// } 
	// req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(configs.ShopId + ":" + string(file))))
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println(err)
	// 	return model.ResponseUKassa{}, err
	// }
	// defer resp.Body.Close()

	// var responseUKassa model.ResponseUKassa
	// err = json.NewDecoder(resp.Body).Decode(&responseUKassa)
	// if err != nil {
	// 	log.Println(err)
	// 	return model.ResponseUKassa{}, err
	// }
	return responseUKassa, nil
}

func Subscribe(payment model.Payment) (model.ResponseUKassa, error) {
	paymentDB, requestUkassa, err := createPaymentUKassa(payment, true)
	if err != nil {
		return  model.ResponseUKassa{} ,err
	}
	requestJson, err := json.Marshal(requestUkassa)
	if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	} 
	reader := bytes.NewReader(requestJson)
	responseUKassa, err := makeRequestUKassa(reader, "POST", paymentDB.UUID)
		if err != nil {
		log.Println(err)
		return model.ResponseUKassa{}, err
	}
	return responseUKassa, nil
}

func requestPaymentStatusAPI(payment model.Payment) (int, string, error) {
	// req, err := http.NewRequest("GET", configs.PaymentAPI + "payments/" + payment.UUID, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return 0, err
	// } 

	// fileDir, _ := os.Getwd()
	// filePath := filepath.Join(fileDir, "API_key")
	// file, err := os.ReadFile(filePath)
	// if err != nil {
	// 	log.Println(err)
	// 	return 0, err
	// } 
	// req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(configs.ShopId + ":" + string(file))))
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println(err)
	// 	return 0, err
	// }
	// defer resp.Body.Close()

	// var responseUKassa model.ResponseUKassa
	// err = json.NewDecoder(resp.Body).Decode(&responseUKassa)

	responseUKassa, err := makeRequestUKassa(nil, "GET", payment.UUID)
	if err != nil {
		log.Println(err)
		return 0, "", err
	}
	switch responseUKassa.Status {
	case "pending":
		return model.PaymentWaitingStatus, responseUKassa.PaymentMethodData.Id, nil
	case "succeeded":
		return model.PaymentSucceededStatus, responseUKassa.PaymentMethodData.Id, nil
	case "canceled":
		return model.PaymentCanceledStatus, responseUKassa.PaymentMethodData.Id, nil
	case "waiting_for_capture":
		return model.PaymentSucceededStatus, responseUKassa.PaymentMethodData.Id, nil
	}

	return 0, "", errors.New("unexpected behaviour")
}

func CheckPaymentStatusAPI(paymentRepository payments.PaymentRepository, payment model.Payment) (model.Payment, error) {
	status, paymentMethodId, err := requestPaymentStatusAPI(payment)
	if err != nil {
		log.Println(err)
		return model.Payment{}, err
	}
	fmt.Println(status)
	payment.PaymentMethodId = paymentMethodId
	payment.Status = uint(status)
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

func MakeCronCheckSubscriptions(paymentRepository payments.PaymentRepository,
	subscriptionRepository subscriptions.SubscriptionRepository) (error) {
		c := cron.New()
		_, err := c.AddFunc("* * * * *", func() {
			fmt.Println("joba3")
			subscriptions, err := subscriptionRepository.GetAllNotFreeSubscriptions()
			if err != nil {
				log.Println(err)
				return
			}
			for _, subscription := range subscriptions {
				lastPayment, err := paymentRepository.GetLastSuccessfulSubscriptionPayment(
					int(subscription.Subscriber_id),
					int(subscription.Creator_id),
				)
				if err != nil {
					log.Println(err)
					return
				}
				payDay, err := time.Parse("2006-01-02", lastPayment.CreatedAt[0:10])
				if (err != nil) {
					log.Println(err)
					return
				}
				payDay.AddDate(0, 1, 0)
				// if (payDay.AddDate(0, 1, 0).After(time.Now())) {
				// 	continue
				// }
				value := strconv.Itoa(int(lastPayment.PaymentInteger))
				if lastPayment.PaymentFractional < 10 {
					value += ".0" + strconv.Itoa(int(lastPayment.PaymentFractional))
				} else {
					value += "." +strconv.Itoa(int(lastPayment.PaymentFractional))
				}
				payment, requestUkassa, err := createAutoPaymentUKassa(
					model.Payment{
						DonaterId: lastPayment.DonaterId,
						CreatorId: lastPayment.CreatorId,
						PaymentInteger: lastPayment.PaymentInteger,
						PaymentFractional: lastPayment.PaymentFractional,
						Currency: lastPayment.Currency,
						Value: value,
						Type: model.PaymentTypeSubscription,
						PaymentMethodId: lastPayment.PaymentMethodId,
					},
				)
				if err != nil {
					log.Println(err)
					return
				}
				requestUkassa.PaymentMethodId = lastPayment.PaymentMethodId
				requestJson, err := json.Marshal(requestUkassa)
				if err != nil {
					log.Println(err)
					return
				} 
				reader := bytes.NewReader(requestJson)
				responseUkassa, err := makeRequestUKassa(reader, "POST", payment.UUID)
				if err != nil {
					log.Println(err)
					return
				}

				switch responseUkassa.Status {
				case "pending":
					payment.Status = model.PaymentWaitingStatus
				case "succeeded":
					payment.Status = model.PaymentSucceededStatus
				case "canceled":
					payment.Status = model.PaymentCanceledStatus
				}
				id, err := paymentRepository.CreateNewPayment(payment)
				if err != nil {
					//think
					log.Println(err)
					return
				}
				payment.Id = id

				cr := cron.New()
				_, err = cr.AddFunc("* * * * *", func() {
					fmt.Println("joba4")
					payment, err := CheckPaymentStatusAPI(paymentRepository, payment)
					if err != nil {
						log.Println(err)
						return
					}
					if payment.Status == model.PaymentSucceededStatus {
						cr.Stop()
						return
					}
					if payment.Status == model.PaymentCanceledStatus {
						err = subscriptionRepository.DeleteSubscription(
							int(subscription.Subscription_level_id), 
							int(subscription.Subscriber_id),
						)
						if err != nil {
							log.Println(err)
							return 
						}
						cr.Stop()
					}
				})
				if err != nil {
					log.Println(err)
					return
				}
				cr.Start()
			}
		})
		if err != nil {
			log.Println(err)
			return err
		}
		c.Start()
		return nil
}