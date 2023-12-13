package services

import (
	"log"
	"time"

	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
)

type CustomerService interface {
	GetListCustomers() ([]*ent.Customer, error)
	GetCustomerByID(ID string) (*ent.Customer, error)
	AddNewCustomer(newCustomerRequest requests.DtoCustomerRequest) (*ent.Customer, error)
	UpdateCustomerData(updateCustomerRequest requests.DtoCustomerRequest, existingCustomer *ent.Customer) error
	DeleteCustomerByID(ID string) error
	GetListDebtorCustomers() ([]*ent.Customer, error)
}

type CustomerServiceImpl struct {
	customerRepo repositories.CustomerRepository
}

func CustomerServiceInit(customerRepo repositories.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		customerRepo: customerRepo,
	}
}

func (cs *CustomerServiceImpl) GetListCustomers() ([]*ent.Customer, error) {

	return cs.customerRepo.ReadAllCustomers()
}

func (cs *CustomerServiceImpl) GetCustomerByID(ID string) (*ent.Customer, error) {

	return cs.customerRepo.ReadCustomerByID(ID)
}

func (cs *CustomerServiceImpl) AddNewCustomer(newCustomerRequest requests.DtoCustomerRequest) (*ent.Customer, error) {
	birthDate, err := helpers.DDmmYYTimeConverterFromStringToTime(newCustomerRequest.BirthDateStr)
	if err != nil {
		log.Fatal("Got an error when parsing string to time birth date data. When registration new customer. ERROR:", err)
	}

	newCustomer := &ent.Customer{
		FullName:           newCustomerRequest.FullName,
		BirthDate:          birthDate,
		RegistrationDate:   time.Now(),
		CustomerPictureURL: newCustomerRequest.CustomerPictureURL,
		Status:             newCustomerRequest.Status,
	}

	return cs.customerRepo.CreateCustomer(newCustomer)
}

func (cs *CustomerServiceImpl) UpdateCustomerData(updateCustomerRequest requests.DtoCustomerRequest, updatedCustomer *ent.Customer) error {

	birthDate, err := helpers.DDmmYYTimeConverterFromStringToTime(updateCustomerRequest.BirthDateStr)
	if err != nil {
		log.Fatal("Got an error when parsing string to time birth date data. When registration new customer. ERROR:", err)
	}

	updatedCustomer.FullName = updateCustomerRequest.FullName
	updatedCustomer.BirthDate = birthDate
	updatedCustomer.CustomerPictureURL = updateCustomerRequest.CustomerPictureURL
	updatedCustomer.Status = updateCustomerRequest.Status

	return cs.customerRepo.UpdateCustomer(updatedCustomer)
}

func (cs *CustomerServiceImpl) DeleteCustomerByID(ID string) error {
	return cs.customerRepo.RemoveCustomerByID(ID)
}

func (cs *CustomerServiceImpl) GetListDebtorCustomers() ([]*ent.Customer, error) {
	return cs.customerRepo.ReadAllDebtors()
}
