package services

import (
	"log"
	"time"

	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/payment"
	"github.com/hmertakyatan/gymapi/repositories"
)

type PaymentService interface {
	ListAllPayments() ([]*ent.Payment, error)
	GetPaymentByID(ID string) (*ent.Payment, error)
	CreatePayment(request requests.DtoPaymentRequest) (*ent.Payment, error)
	UpdatePaymentByID(ID string, request requests.DtoPaymentRequest) (*ent.Payment, error)
	DeletePaymentByID(ID string) error
	ListIncomePayments() ([]*ent.Payment, error)
	ListExpensePayments() ([]*ent.Payment, error)
}

type PaymentServiceImpl struct {
	paymentRepo repositories.PaymentRepository
}

func PaymentServiceInit(paymentRepo repositories.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepo: paymentRepo,
	}
}

func (ps *PaymentServiceImpl) ListAllPayments() ([]*ent.Payment, error) {

	return ps.paymentRepo.ReadAllPayments()
}

func (ps *PaymentServiceImpl) GetPaymentByID(ID string) (*ent.Payment, error) {

	return ps.paymentRepo.ReadPaymentByID(ID)
}

func (ps *PaymentServiceImpl) CreatePayment(request requests.DtoPaymentRequest) (*ent.Payment, error) {

	newPayment := &ent.Payment{
		Type:        payment.Type(request.Type),
		Description: request.Description,
		Amount:      request.Amount,
		DateTime:    time.Now(),
	}

	return ps.paymentRepo.CreatePayment(newPayment)
}

func (ps *PaymentServiceImpl) UpdatePaymentByID(ID string, request requests.DtoPaymentRequest) (*ent.Payment, error) {

	updatedPayment, err := ps.paymentRepo.ReadPaymentByID(ID)
	if err != nil {
		log.Fatal("Got an error getting payment datas: ERROR :", err)
	}

	updatedPayment.Type = payment.Type(request.Type)
	updatedPayment.Description = request.Description
	updatedPayment.DateTime = time.Now()
	updatedPayment.Amount = request.Amount

	return ps.paymentRepo.UpdatePayment(updatedPayment)

}

func (ps *PaymentServiceImpl) DeletePaymentByID(ID string) error {

	return ps.paymentRepo.RemovePaymentByID(ID)

}

func (ps *PaymentServiceImpl) ListIncomePayments() ([]*ent.Payment, error) {

	return ps.paymentRepo.ReadAllIncomePayments()
}

func (ps *PaymentServiceImpl) ListExpensePayments() ([]*ent.Payment, error) {

	return ps.paymentRepo.ReadAllExpensePayments()
}
