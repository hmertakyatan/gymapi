package services

import (
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/repositories"
)

type MembershipTypeService interface {
	GetListMembershipTypes() ([]*ent.Membership_type, error)
	GetMembershipTypeByID(ID string) (*ent.Membership_type, error)
	AddNewMembershipType(newMembershipTypeRequest requests.DtoMembershipTypeRequest) (*ent.Membership_type, error)
	UpdateMembershipTypeData(updateMembershipTypeRequest requests.DtoMembershipTypeRequest, existingMembershipType *ent.Membership_type) error
	DeleteMembershipTypeByID(ID string) error
}

type MembershipTypeServiceImpl struct {
	membershipTypeRepo repositories.MembershipTypeRepository
}

func MembershipTypeServiceInit(membershipTypeRepo repositories.MembershipTypeRepository) MembershipTypeService {
	return &MembershipTypeServiceImpl{
		membershipTypeRepo: membershipTypeRepo,
	}
}

func (mts *MembershipTypeServiceImpl) GetListMembershipTypes() ([]*ent.Membership_type, error) {

	return mts.membershipTypeRepo.ReadAllMembershipTypes()
}

func (mts *MembershipTypeServiceImpl) GetMembershipTypeByID(ID string) (*ent.Membership_type, error) {
	return mts.membershipTypeRepo.ReadMembershipTypeByID(ID)

}

func (mts *MembershipTypeServiceImpl) AddNewMembershipType(newMembershipTypeRequest requests.DtoMembershipTypeRequest) (*ent.Membership_type, error) {

	newMembershipType := &ent.Membership_type{
		Name:            newMembershipTypeRequest.Name,
		Description:     newMembershipTypeRequest.Description,
		MembershipMonth: newMembershipTypeRequest.MembershipMonth,
		Price:           newMembershipTypeRequest.Price,
	}

	return mts.membershipTypeRepo.CreateMembershipType(newMembershipType)

}

func (mts *MembershipTypeServiceImpl) UpdateMembershipTypeData(updateMembershipTypeRequest requests.DtoMembershipTypeRequest, updatedMembershipType *ent.Membership_type) error {

	updatedMembershipType.Name = updateMembershipTypeRequest.Name
	updatedMembershipType.Description = updateMembershipTypeRequest.Description
	updatedMembershipType.MembershipMonth = updateMembershipTypeRequest.MembershipMonth
	updatedMembershipType.Price = updateMembershipTypeRequest.Price

	return mts.membershipTypeRepo.UpdateMembershipType(updatedMembershipType)

}

func (mts *MembershipTypeServiceImpl) DeleteMembershipTypeByID(ID string) error {

	return mts.membershipTypeRepo.RemoveMembershipTypeByID(ID)

}
