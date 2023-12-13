package services

import (
	"log"
	"time"

	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
)

type PersonnelService interface {
	GetListPersonnels() ([]*ent.Personnel, error)
	GetPersonnelByID(ID string) (*ent.Personnel, error)
	AddNewPersonnel(PersonnelRequest requests.DtoPersonnelRequest) (*ent.Personnel, error)
	UpdatePersonnelData(ID string, PersonnelUpdateRequest requests.DtoPersonnelRequest) (*ent.Personnel, error)
	DeletePersonnelByID(ID string) error
}

type PersonnelServiceImpl struct {
	personnelRepo repositories.PersonnelRepository
}

func PersonnelServiceInit(personnelRepo repositories.PersonnelRepository) PersonnelService {

	return &PersonnelServiceImpl{
		personnelRepo: personnelRepo,
	}
}

func (p *PersonnelServiceImpl) GetListPersonnels() ([]*ent.Personnel, error) {
	personnels, err := p.personnelRepo.ReadAllPersonnel()

	if err != nil {
		log.Fatal("Got an error when getting list personnels. ERROR: ", err)
	}

	return personnels, nil
}

func (p *PersonnelServiceImpl) GetPersonnelByID(ID string) (*ent.Personnel, error) {
	personnel, err := p.personnelRepo.ReadPersonnelByID(ID)
	if err != nil {
		log.Fatal("Got an error when getting personel by id. ERROR: ", err)
	}
	return personnel, nil
}

func (p *PersonnelServiceImpl) AddNewPersonnel(PersonnelRequest requests.DtoPersonnelRequest) (*ent.Personnel, error) {
	birthDate, err := helpers.DDmmYYTimeConverterFromStringToTime(PersonnelRequest.BirthDateStr)
	if err != nil {
		log.Fatal("Got an error when parsing string to time birth date data. When registration new customer. ERROR:", err)
	}
	startDate, err := helpers.DDmmYYTimeConverterFromStringToTime(PersonnelRequest.StartDate)
	if err != nil {
		log.Fatal("Got an error when parsing string to time birth date data. When registration new customer. ERROR:", err)
	}

	newPersonnel := &ent.Personnel{
		Name:        PersonnelRequest.Name,
		Description: PersonnelRequest.Description,
		BithDate:    birthDate,
		StartDate:   startDate,
		Salary:      PersonnelRequest.Salary,
		Status:      PersonnelRequest.Status,
	}

	return p.personnelRepo.CreatePersonnel(newPersonnel)
}

func (p *PersonnelServiceImpl) UpdatePersonnelData(ID string, PersonnelUpdateRequest requests.DtoPersonnelRequest) (*ent.Personnel, error) {
	personnel, err := p.personnelRepo.ReadPersonnelByID(ID)
	if err != nil {
		log.Fatal("Got an error when reading personnel by id. ERROR :", err)
	}

	birthDate, err := time.Parse("02/01/2006", PersonnelUpdateRequest.BirthDateStr)
	if err != nil {
		log.Fatal("Got error when parsing personnel's birth date data. ERROR :", err)
	}
	startDate, err := time.Parse("02/01/2006", PersonnelUpdateRequest.StartDate)
	if err != nil {
		log.Fatal("Got error when parsing personnel's start date data. ERROR :", err)
	}
	personnel.Name = PersonnelUpdateRequest.Name
	personnel.Salary = PersonnelUpdateRequest.Salary
	personnel.StartDate = startDate
	personnel.BithDate = birthDate
	personnel.Status = PersonnelUpdateRequest.Status

	return p.personnelRepo.UpdatePersonnel(personnel)

}

func (p *PersonnelServiceImpl) DeletePersonnelByID(ID string) error {
	return p.personnelRepo.RemovePersonnelByID(ID)
}
