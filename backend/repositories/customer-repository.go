package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/customer"
	"github.com/hmertakyatan/gymapi/ent/membership"
)

type CustomerRepository interface {
	ReadAllCustomers() ([]*ent.Customer, error)
	ReadCustomerByID(ID string) (*ent.Customer, error)
	CreateCustomer(customer *ent.Customer) (*ent.Customer, error)
	UpdateCustomer(customer *ent.Customer) error
	RemoveCustomerByID(ID string) error
	ReadAllDebtors() ([]*ent.Customer, error)
}

type CustomerRepositoryImpl struct {
	client *ent.Client
}

func CustomerRepositoryInit(client *ent.Client) CustomerRepository {
	return &CustomerRepositoryImpl{client: client}
}

func (r *CustomerRepositoryImpl) ReadAllCustomers() ([]*ent.Customer, error) {
	customers, err := r.client.Customer.Query().WithMembership().All(context.Background())
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepositoryImpl) ReadCustomerByID(ID string) (*ent.Customer, error) {
	customerID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatal("Got an error when parsing customer id. ERROR: ", err)
		return nil, err
	}
	customer, err := r.client.Customer.Query().Where(customer.ID(customerID)).WithMembership().Only(context.Background())
	if err != nil {
		log.Fatal("Got an error when finding customer by id. ERROR: ")
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepositoryImpl) CreateCustomer(customer *ent.Customer) (*ent.Customer, error) {
	createdCustomer, err := r.client.Customer.Create().
		SetFullName(customer.FullName).
		SetBirthDate(customer.BirthDate).
		SetRegistrationDate(customer.RegistrationDate).
		SetCustomerPictureURL(customer.CustomerPictureURL).
		SetStatus(customer.Status).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}

func (r *CustomerRepositoryImpl) UpdateCustomer(changedcustomer *ent.Customer) error {
	err := r.client.Customer.UpdateOneID(changedcustomer.ID).
		SetFullName(changedcustomer.FullName).
		SetBirthDate(changedcustomer.BirthDate).
		SetCustomerPictureURL(changedcustomer.CustomerPictureURL).
		SetStatus(changedcustomer.Status).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepositoryImpl) RemoveCustomerByID(ID string) error {
	customerID, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	err = r.client.Customer.DeleteOneID(customerID).Exec(context.Background())
	return err
}

func (r *CustomerRepositoryImpl) ReadAllDebtors() ([]*ent.Customer, error) {
	debtors, err := r.client.Customer.Query().Where(customer.HasMembershipWith(membership.AmountRemainingGT(0))).WithMembership().All(context.Background())
	if err != nil {
		return nil, err
	}
	return debtors, nil
}
