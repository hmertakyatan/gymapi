package repositories

import (
	"context"
	"log"

	"github.com/hmertakyatan/gymapi/ent"
)

type PersonalTrainingRepository interface {
	CreateRelationBetweenPersonnelAndCustomer(personnel *ent.Personnel, customer *ent.Customer, registartion *ent.PersonnelRelCustomer) (*ent.PersonnelRelCustomer, error)
}

type PersonalTrainingRepositoryImpl struct {
	client *ent.Client
}

func PersonalTrainingRepositoryInit(client *ent.Client) PersonalTrainingRepository {
	return &PersonalTrainingRepositoryImpl{client: client}
}

func (r *PersonalTrainingRepositoryImpl) CreateRelationBetweenPersonnelAndCustomer(personnel *ent.Personnel, customer *ent.Customer, registration *ent.PersonnelRelCustomer) (*ent.PersonnelRelCustomer, error) {
	createdRelationBetweenPersonnelAndCustomer, err := r.client.PersonnelRelCustomer.Create().
		SetDescription(registration.Description).
		SetPrice(registration.Price).
		SetCustomer(customer).
		SetPersonnel(personnel).
		Save(context.Background())

	if err != nil {
		log.Fatal("Got an error when saving relation beetwen personnel and customer. ERROR :", err)
	}

	return createdRelationBetweenPersonnelAndCustomer, nil
}
