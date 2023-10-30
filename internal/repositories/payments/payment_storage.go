package payments

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertPaymentSQL(payment model.Payment) squirrel.InsertBuilder {
	return squirrel.Insert(configs.PaymentTable).
		Columns("payment_integer", "payment_fractional", "status", "donater_id", "creator_id").
		Values(payment.PaymentInteger, payment.PaymentFractional, payment.Status, payment.DonaterId, payment.DonaterId).
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
	return squirrel.Select("*").
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
		Where(squirrel.Eq{"uuid": payment.UUID}).
		PlaceholderFormat(squirrel.Dollar)
}

type PaymentStorage struct {
	db *sql.DB
}

func CreatePaymentStorage(db *sql.DB) PaymentRepository {
	return &PaymentStorage{
		db: db,
	}
}

func (storage *PaymentStorage) CreateNewPayment(payment model.Payment) (int, error) {
	var paymentId int
	err := InsertPaymentSQL(payment).RunWith(storage.db).QueryRow().Scan(paymentId)
	if err != nil {
		return 0, err
	}
	return paymentId, nil
}

func (storage *PaymentStorage) DeletePayment(id uint) error {
	_, err := DeletePaymentSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
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

func (storage *PaymentStorage) ChangePayment(payment model.Payment) error {
	_, err := UpdatePaymentSQL(payment).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
}

func (storage *PaymentStorage) GetPaymentsByAuthorId(authorId uint) ([]model.Payment, error) {
	rows, err := SelectPaymentsByAuthorIdSQL(authorId).RunWith(storage.db).Query()
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
