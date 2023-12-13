package services

import (
	"log"
	"time"

	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
)

type MembershipService interface {
	GetListMemberships() ([]*ent.Membership, error)
	GetMembershipByID(ID string) (*ent.Membership, error)
	AddNewMembership(newMembershipRequest requests.DtoMembershipRequest) (*ent.Membership, error)
	UpdateMembershipData(updatedMembershipRequest requests.DtoMembershipRequest, existingMembership *ent.Membership) error
	DeleteMembershipByID(ID string) error
}

type MembershipServiceImpl struct {
	membershipRepo     repositories.MembershipRepository
	membershipTypeRepo repositories.MembershipTypeRepository
	customerRepo       repositories.CustomerRepository
}

func MembershipServiceInit(membershipTypeRepo repositories.MembershipTypeRepository, customerRepo repositories.CustomerRepository,
	membershipRepo repositories.MembershipRepository) MembershipService {
	return &MembershipServiceImpl{
		membershipRepo:     membershipRepo,
		membershipTypeRepo: membershipTypeRepo,
		customerRepo:       customerRepo,
	}
}

func (ms *MembershipServiceImpl) GetListMemberships() ([]*ent.Membership, error) {
	return ms.membershipRepo.ReadAllMemberships()
}

func (ms *MembershipServiceImpl) GetMembershipByID(ID string) (*ent.Membership, error) {
	return ms.membershipRepo.ReadMembershipByID(ID)
}

func (ms *MembershipServiceImpl) AddNewMembership(newMembershipRequest requests.DtoMembershipRequest) (*ent.Membership, error) {

	selectedCustomer, err := ms.customerRepo.ReadCustomerByID(newMembershipRequest.CustomerID)
	if err != nil {
		log.Fatal("Got an error when getting customer datas. ERROR: ", err)
	}

	startDate, err := helpers.DDmmYYTimeConverterFromStringToTime(newMembershipRequest.StartDateStr)
	if err != nil {
		log.Fatal("Got an error when parsing birth date data. ERROR: ", err)
	}
	membershipTypeDatas, err := ms.membershipTypeRepo.ReadMembershipTypeByID(newMembershipRequest.MembershipTypeID)
	if err != nil {
		log.Fatal("Got an error when getting membership type datas. ERROR :", err)
	}
	endDate, err := helpers.CalculateEndDate(startDate, membershipTypeDatas.MembershipMonth)
	if err != nil {
		log.Fatal("Got an error when calculating end date. ERROR :", err)
	}

	calculatedAmountRemaining := membershipTypeDatas.Price - newMembershipRequest.AmountPaid

	newMembership := &ent.Membership{
		OperationDate:   time.Now(),
		StartDate:       startDate,
		EndDate:         endDate,
		AmountPaid:      newMembershipRequest.AmountPaid,
		AmountRemaining: calculatedAmountRemaining,
	}

	return ms.membershipRepo.CreateMembership(newMembership, selectedCustomer)
}

func (ms *MembershipServiceImpl) UpdateMembershipData(updateMembershipRequest requests.DtoMembershipRequest, updatedMembership *ent.Membership) error {
	startDate, err := helpers.DDmmYYTimeConverterFromStringToTime(updateMembershipRequest.StartDateStr)
	if err != nil {
		log.Fatal("Got an error when parsing startdate data. ERROR: ", err)
	}

	memebershipTypeInfo, err := ms.membershipTypeRepo.ReadMembershipTypeByID(updateMembershipRequest.MembershipTypeID)
	if err != nil {
		log.Fatal("Got an error when reading membership type datas. ERROR :", err)
	}

	endDate, err := helpers.CalculateEndDate(startDate, memebershipTypeInfo.MembershipMonth)
	if err != nil {
		log.Fatal("Got an error when calculating end date. ERROR: ", err)
	}
	amountRemaining := updateMembershipRequest.AmountPaid - memebershipTypeInfo.Price

	updatedMembership.OperationDate = time.Now()
	updatedMembership.AmountPaid = updateMembershipRequest.AmountPaid
	updatedMembership.StartDate = startDate
	updatedMembership.EndDate = endDate
	updatedMembership.AmountRemaining = amountRemaining

	return ms.membershipRepo.UpdateMembership(updatedMembership)
}

func (ms *MembershipServiceImpl) DeleteMembershipByID(ID string) error {
	return ms.membershipRepo.RemoveMembershipByID(ID)
}
