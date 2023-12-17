package payments

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertPaymentSQL(payment model.Payment) squirrel.InsertBuilder {
	return squirrel.Insert(configs.PaymentTable).
		Columns("payment_integer", "payment_fractional", "status", "donater_id", "creator_id").
		Values(payment.PaymentInteger, payment.PaymentFractional, payment.Status, payment.DonaterId, payment.CreatorId).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeletePaymentSQL(paymentId uint) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.PaymentTable).
		Where(squirrel.Eq{"id": paymentId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectPaymentByUUIDSQL(paymentId string) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PaymentTable).
		Where(squirrel.Eq{"uuid": paymentId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectPaymentsByAuthorIdSQL(authorId uint) squirrel.SelectBuilder {
	return squirrel.Select(
		"coalesce(sum(p.payment_integer) FILTER  (WHERE p.status = 2), 0) as payment_integer",
		"coalesce(sum(p.payment_fractional) FILTER (WHERE p.status = 2), 0) AS payment_fractional").
		From(configs.PaymentTable + " p").
		Where(squirrel.Eq{"p.creator_id": authorId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectPaymentsByUserIdSQL(userId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PaymentTable + " p").
		Where(squirrel.Eq{"p.donater_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func UpdatePaymentSQL(payment model.Payment) squirrel.UpdateBuilder {
	return squirrel.Update(configs.PaymentTable).
		SetMap(map[string]interface{}{
			"status": payment.Status,
		}).
		// Where(squirrel.Eq{"uuid": payment.UUID}).
		Where(squirrel.Eq{"id": payment.Id}).
		PlaceholderFormat(squirrel.Dollar)
}

type PaymentStorage struct {
	UnimplementedPaymentsServiceServer
	db *sql.DB
}

type PaymentManager struct {
	Client PaymentsServiceClient
}

func (manager *PaymentManager) CreateNewPayment(payment model.Payment) (int, error) {
	i, err := manager.Client.CreateNewPaymentCtx(context.Background(), &PaymentGRPC{
		Id:             int32(payment.Id),
		Uuid:           payment.UUID,
		PaymentInteger: uint32(payment.PaymentInteger),
		Status:         uint32(payment.Status),
		DonaterId:      uint32(payment.DonaterId),
		CreatorId:      uint32(payment.CreatorId),
		Currency:       payment.Currency,
		Value:          payment.Value,
	})

	return int(i.I), err
}

func (manager *PaymentManager) DeletePayment(id uint) error {
	_, err := manager.Client.DeletePaymentCtx(context.Background(), &UInt{Id: uint32(id)})

	return err
}

func (manager *PaymentManager) GetPayment(uuid string) (model.Payment, error) {
	payment, err := manager.Client.GetPaymentCtx(context.Background(), &UUid{Uuid: uuid})
	return model.Payment{
		Id:             int(payment.Id),
		UUID:           payment.Uuid,
		PaymentInteger: uint(payment.PaymentInteger),
		Status:         uint(payment.Status),
		DonaterId:      uint(payment.DonaterId),
		CreatorId:      uint(payment.CreatorId),
		Currency:       payment.Currency,
		Value:          payment.Value,
	}, err
}

func (manager *PaymentManager) ChangePayment(payment model.Payment) error {
	_, err := manager.Client.ChangePaymentCtx(context.Background(), &PaymentGRPC{
		Id:             int32(payment.Id),
		Uuid:           payment.UUID,
		PaymentInteger: uint32(payment.PaymentInteger),
		Status:         uint32(payment.Status),
		DonaterId:      uint32(payment.DonaterId),
		CreatorId:      uint32(payment.CreatorId),
		Currency:       payment.Currency,
		Value:          payment.Value,
	})

	return err
}

func (manager *PaymentManager) GetPaymentsByAuthorId(authorId uint) ([]model.Payment, error) {
	paymentsGRPC, err := manager.Client.GetPaymentsByAuthorIdCtx(context.Background(), &UInt{Id: uint32(authorId)})

	if err != nil {
		return []model.Payment{}, err
	}

	var payments []model.Payment
	for _, payment := range paymentsGRPC.Payments {
		payments = append(payments, model.Payment{
			Id:             int(payment.Id),
			UUID:           payment.Uuid,
			PaymentInteger: uint(payment.PaymentInteger),
			PaymentFractional: uint(payment.PaymentFractional),
			Status:         uint(payment.Status),
			DonaterId:      uint(payment.DonaterId),
			CreatorId:      uint(payment.CreatorId),
			Currency:       payment.Currency,
			Value:          payment.Value,
		})
	}
	return payments, nil
}

func (manager *PaymentManager) GetPaymentsByUserId(userId uint) ([]model.Payment, error) {
	paymentsMap, err := manager.Client.GetPaymentsByUserIdCtx(context.Background(), &UInt{Id: uint32(userId)})
	
	if err != nil {
		return nil, err
	}

	var payments []model.Payment

	for _, payment := range paymentsMap.Payments {
		payments = append(payments, model.Payment{
			Id:             int(payment.Id),
			UUID:           payment.Uuid,
			PaymentInteger: uint(payment.PaymentInteger),
			Status:         uint(payment.Status),
			DonaterId:      uint(payment.DonaterId),
			CreatorId:      uint(payment.CreatorId),
			Currency:       payment.Currency,
			Value:          payment.Value,
		})
	}

	return payments, nil
}

// func CreatePaymentStorage(db *sql.DB) PaymentRepository {
// 	return &PaymentStorage{
// 		db: db,
// 	}
// }

func CreatePaymentStore(db *sql.DB) *PaymentStorage {
	return &PaymentStorage{
		db: db,
	}
}

func (storage *PaymentStorage) CreateNewPayment(payment model.Payment) (int, error) {
	var paymentId int
	err := InsertPaymentSQL(payment).RunWith(storage.db).QueryRow().Scan(&paymentId)
	if err != nil {
		return 0, err
	}
	return paymentId, nil
}

func (storage *PaymentStorage) CreateNewPaymentCtx(ctx context.Context, payment *PaymentGRPC) (*Int, error) {
	var paymentId int
	err := InsertPaymentSQL(model.Payment{
		Id:             int(payment.Id),
		UUID:           payment.Uuid,
		PaymentInteger: uint(payment.PaymentInteger),
		Status:         uint(payment.Status),
		DonaterId:      uint(payment.DonaterId),
		CreatorId:      uint(payment.CreatorId),
		Currency:       payment.Currency,
		Value:          payment.Value,
	}).RunWith(storage.db).QueryRow().Scan(&paymentId)
	if err != nil {
		return &Int{
			I: 0,
		}, err
	}
	return &Int{
		I: payment.Id,
	}, nil
}

func (storage *PaymentStorage) DeletePayment(id uint) error {
	rows, err := DeletePaymentSQL(id).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func (storage *PaymentStorage) DeletePaymentCtx(ctx context.Context, id *UInt) (*Nothing, error) {
	rows, err := DeletePaymentSQL(uint(id.Id)).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return &Nothing{}, err
	}
	return &Nothing{}, nil
}

func (storage *PaymentStorage) GetPayment(uuid string) (model.Payment, error) {
	rows, err := SelectPaymentByUUIDSQL(uuid).RunWith(storage.db).Query()
	if err != nil {
		return model.Payment{}, err
	}
	var payments []model.Payment
	if err = dbscan.ScanAll(&payments, rows); err != nil {
		return model.Payment{}, err
	}
	if len(payments) > 0 {
		return payments[0], nil
	}
	return model.Payment{}, nil
}

func (storage *PaymentStorage) GetPaymentCtx(ctx context.Context, uuid *UUid) (*PaymentGRPC, error) {
	rows, err := SelectPaymentByUUIDSQL(uuid.Uuid).RunWith(storage.db).Query()
	if err != nil {
		return &PaymentGRPC{}, err
	}
	var payments []model.Payment
	if err = dbscan.ScanAll(&payments, rows); err != nil {
		return &PaymentGRPC{}, err
	}
	if len(payments) > 0 {
		return &PaymentGRPC{
			Id:             int32(payments[0].Id),
			Uuid:           payments[0].UUID,
			PaymentInteger: uint32(payments[0].PaymentInteger),
			Status:         uint32(payments[0].Status),
			DonaterId:      uint32(payments[0].DonaterId),
			CreatorId:      uint32(payments[0].CreatorId),
			Currency:       payments[0].Currency,
			Value:          payments[0].Value,
		}, nil
	}
	return &PaymentGRPC{}, nil
}

func (storage *PaymentStorage) ChangePayment(payment model.Payment) error {
	rows, err := UpdatePaymentSQL(payment).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func (storage *PaymentStorage) ChangePaymentCtx(ctx context.Context, payment *PaymentGRPC) (*Nothing, error) {
	fmt.Println(payment.Status)
	rows, err := UpdatePaymentSQL(model.Payment{
		Id:             int(payment.Id),
		UUID:           payment.Uuid,
		PaymentInteger: uint(payment.PaymentInteger),
		Status:         uint(payment.Status),
		DonaterId:      uint(payment.DonaterId),
		CreatorId:      uint(payment.CreatorId),
		Currency:       payment.Currency,
		Value:          payment.Value,
	}).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return &Nothing{}, err
	}
	return &Nothing{}, nil
}

func (storage *PaymentStorage) GetPaymentsByAuthorId(authorId uint) (model.Payment, error) {
	rows, err := SelectPaymentsByAuthorIdSQL(authorId).RunWith(storage.db).Query()
	if err != nil {
		return model.Payment{}, err
	}
	var payments []model.Payment
	err = dbscan.ScanAll(&payments, rows)
	if err != nil {
		return model.Payment{}, err
	}
	if len(payments) > 0 {
		return payments[0], nil
	}
	return model.Payment{}, nil
}

func (storage *PaymentStorage) GetPaymentsByAuthorIdCtx(ctx context.Context, authorID *UInt) (*PaymentsGRPC, error) {
	rows, err := SelectPaymentsByAuthorIdSQL(uint(authorID.Id)).RunWith(storage.db).Query()
	if err != nil {
		return &PaymentsGRPC{}, err
	}
	var payments []model.Payment
	err = dbscan.ScanAll(&payments, rows)
	if err != nil {
		return &PaymentsGRPC{}, err
	}
	paymetsMap := &PaymentsGRPC{}
	for i, payment := range payments {
		paymetsMap.Payments[int32(i)] = &PaymentGRPC{
			Id:             int32(payment.Id),
			Uuid:           payment.UUID,
			PaymentInteger: uint32(payment.PaymentInteger),
			PaymentFractional: uint32(payment.PaymentFractional),
			Status:         uint32(payment.Status),
			DonaterId:      uint32(payment.DonaterId),
			CreatorId:      uint32(payment.CreatorId),
			Currency:       payment.Currency,
			Value:          payment.Value,
		}
	}
	return paymetsMap, nil
}

func (storage *PaymentStorage) GetPaymentsByUserId(userId uint) ([]model.Payment, error) {
	rows, err := SelectPaymentsByUserIdSQL(userId).RunWith(storage.db).Query()
	if err != nil {
		return []model.Payment{}, err
	}
	var payments []model.Payment
	err = dbscan.ScanAll(&payments, rows)
	if err != nil {
		return []model.Payment{}, err
	}
	return payments, nil
}

func (storage *PaymentStorage) GetPaymentsByUserIdCtx(ctx context.Context, userId *UInt) (*PaymentsGRPC, error) {
	rows, err := SelectPaymentsByUserIdSQL(uint(userId.Id)).RunWith(storage.db).Query()
	if err != nil {
		return &PaymentsGRPC{}, err
	}
	var payments []model.Payment
	err = dbscan.ScanAll(&payments, rows)
	if err != nil {
		return &PaymentsGRPC{}, err
	}

	var paymentsMap *PaymentsGRPC
	for i, payment := range payments {
		paymentsMap.Payments[int32(i)] = &PaymentGRPC{
			Id:             int32(payment.Id),
			Uuid:           payment.UUID,
			PaymentInteger: uint32(payment.PaymentInteger),
			Status:         uint32(payment.Status),
			DonaterId:      uint32(payment.DonaterId),
			CreatorId:      uint32(payment.CreatorId),
			Currency:       payment.Currency,
			Value:          payment.Value,
		}
	}

	return paymentsMap, nil
}