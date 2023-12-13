// payment_repository.go
package repositories

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/payment"
)

type PaymentRepository interface {
	ReadAllPayments() ([]*ent.Payment, error)
	ReadPaymentByID(ID string) (*ent.Payment, error)
	CreatePayment(payment *ent.Payment) (*ent.Payment, error)
	UpdatePayment(payment *ent.Payment) (*ent.Payment, error)
	RemovePaymentByID(ID string) error
	ReadAllIncomePayments() ([]*ent.Payment, error)
	ReadAllExpensePayments() ([]*ent.Payment, error)
}

type PaymentRepositoryImpl struct {
	client *ent.Client
}

func PaymentRepositoryInit(client *ent.Client) PaymentRepository {
	return &PaymentRepositoryImpl{client: client}
}

func (r *PaymentRepositoryImpl) ReadAllPayments() ([]*ent.Payment, error) {
	payments, err := r.client.Payment.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *PaymentRepositoryImpl) ReadPaymentByID(ID string) (*ent.Payment, error) {
	paymentID, err := uuid.Parse(ID)
	if err != nil {
		return nil, err
	}
	payment, err := r.client.Payment.Query().Where(payment.ID(paymentID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepositoryImpl) CreatePayment(payment *ent.Payment) (*ent.Payment, error) {
	createdPayment, err := r.client.Payment.Create().
		SetType(payment.Type).
		SetDescription(payment.Description).
		SetDateTime(payment.DateTime).
		SetAmount(payment.Amount).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return createdPayment, nil
}

func (r *PaymentRepositoryImpl) UpdatePayment(changedpayment *ent.Payment) (*ent.Payment, error) {
	updatedPayment, err := r.client.Payment.UpdateOneID(changedpayment.ID).
		SetType(changedpayment.Type).
		SetAmount(changedpayment.Amount).
		SetDescription(changedpayment.Description).
		SetDateTime(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedPayment, nil
}

func (r *PaymentRepositoryImpl) RemovePaymentByID(ID string) error {
	paymentID, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	err = r.client.Payment.DeleteOneID(paymentID).Exec(context.Background())
	return err
}

func (r *PaymentRepositoryImpl) ReadAllIncomePayments() ([]*ent.Payment, error) {
	incomePayments, err := r.client.Payment.Query().
		Where(payment.TypeEQ("in")).
		All(context.Background())

	if err != nil {
		log.Fatal("Got an error when getting income payment datas from database. ERROR: ", err)
	}
	return incomePayments, nil
}

func (r *PaymentRepositoryImpl) ReadAllExpensePayments() ([]*ent.Payment, error) {
	expensePayments, err := r.client.Payment.Query().
		Where(payment.TypeEQ("out")).
		All(context.Background())

	if err != nil {
		log.Fatal("Got an error when getting expense payment datas from database. ERROR: ", err)
	}

	return expensePayments, nil
}
