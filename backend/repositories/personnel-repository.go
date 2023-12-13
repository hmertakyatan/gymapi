package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/personnel"
)

type PersonnelRepository interface {
	ReadAllPersonnel() ([]*ent.Personnel, error)
	ReadPersonnelByID(ID string) (*ent.Personnel, error)
	CreatePersonnel(personnel *ent.Personnel) (*ent.Personnel, error)
	UpdatePersonnel(personnel *ent.Personnel) (*ent.Personnel, error)
	RemovePersonnelByID(ID string) error
}

type PersonnelRepositoryImpl struct {
	client *ent.Client
}

func PersonnelRepositoryInit(client *ent.Client) PersonnelRepository {
	return &PersonnelRepositoryImpl{client: client}
}

func (r *PersonnelRepositoryImpl) ReadAllPersonnel() ([]*ent.Personnel, error) {
	personnels, err := r.client.Personnel.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return personnels, nil
}

func (r *PersonnelRepositoryImpl) ReadPersonnelByID(ID string) (*ent.Personnel, error) {
	personnelID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatal("Got an error when parsing personnel id", err)
	}

	personnel, err := r.client.Personnel.Query().Where(personnel.ID(personnelID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return personnel, nil
}

func (r *PersonnelRepositoryImpl) CreatePersonnel(personnel *ent.Personnel) (*ent.Personnel, error) {
	createdPersonnel, err := r.client.Personnel.Create().
		SetName(personnel.Name).
		SetDescription(personnel.Description).
		SetSalary(personnel.Salary).
		SetBithDate(personnel.BithDate).
		Save(context.Background())

	if err != nil {
		return nil, err
	}
	return createdPersonnel, nil
}

func (r *PersonnelRepositoryImpl) UpdatePersonnel(personnel *ent.Personnel) (*ent.Personnel, error) {
	updatedPersonnel, err := r.client.Personnel.UpdateOneID(personnel.ID).
		SetName(personnel.Name).
		SetDescription(personnel.Description).
		SetBithDate(personnel.BithDate).
		SetSalary(personnel.Salary).
		SetStatus(personnel.Status).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return updatedPersonnel, nil
}

func (r *PersonnelRepositoryImpl) RemovePersonnelByID(ID string) error {
	personnelID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatal("Got an error when parsing personnel id: ERROR:", err)
	}
	err = r.client.Personnel.DeleteOneID(personnelID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
