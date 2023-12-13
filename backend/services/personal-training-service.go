package services

import (
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/repositories"
)

type PersonalTrainingService interface {
	AddNewPersonalTrainingRegistration(personnel *ent.Personnel, customer *ent.Customer, newPTRegistrationRequest requests.DtoPersonalTrainingRegistartionRequest) (*ent.PersonnelRelCustomer, error)
}

type PersonalTrainingServiceImpl struct {
	personnelTrainingRepo repositories.PersonalTrainingRepository
}

func PersonalTrainingServiceInit(personnelTrainingRepo repositories.PersonalTrainingRepository) PersonalTrainingService {
	return &PersonalTrainingServiceImpl{
		personnelTrainingRepo: personnelTrainingRepo,
	}
}

func (pts *PersonalTrainingServiceImpl) AddNewPersonalTrainingRegistration(personnel *ent.Personnel, customer *ent.Customer, newPTRegistrationRequest requests.DtoPersonalTrainingRegistartionRequest) (*ent.PersonnelRelCustomer, error) {
	newPTRegistration := &ent.PersonnelRelCustomer{
		Description: newPTRegistrationRequest.Description,
		Price:       newPTRegistrationRequest.Price,
	}
	return pts.personnelTrainingRepo.CreateRelationBetweenPersonnelAndCustomer(personnel, customer, newPTRegistration)
}
