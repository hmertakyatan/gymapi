package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/membership_type"
)

type MembershipTypeRepository interface {
	ReadAllMembershipTypes() ([]*ent.Membership_type, error)
	ReadMembershipTypeByID(ID string) (*ent.Membership_type, error)
	CreateMembershipType(membershipType *ent.Membership_type) (*ent.Membership_type, error)
	UpdateMembershipType(membershipType *ent.Membership_type) error
	RemoveMembershipTypeByID(ID string) error
}

type MembershipTypeRepositoryImpl struct {
	client *ent.Client
}

func MembershipTypeRepositoryInit(client *ent.Client) MembershipTypeRepository {
	return &MembershipTypeRepositoryImpl{client: client}
}

func (r *MembershipTypeRepositoryImpl) ReadAllMembershipTypes() ([]*ent.Membership_type, error) {
	membershipTypes, err := r.client.Membership_type.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return membershipTypes, nil
}

func (r *MembershipTypeRepositoryImpl) ReadMembershipTypeByID(ID string) (*ent.Membership_type, error) {
	membershipTypeID, err := uuid.Parse(ID)
	if err != nil {
		return nil, err
	}
	membershipType, err := r.client.Membership_type.Query().Where(membership_type.ID(membershipTypeID)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return membershipType, nil
}

func (r *MembershipTypeRepositoryImpl) CreateMembershipType(membershipType *ent.Membership_type) (*ent.Membership_type, error) {
	createdMembershipType, err := r.client.Membership_type.Create().
		SetName(membershipType.Name).
		SetPrice(membershipType.Price).
		SetDescription(membershipType.Description).
		SetMembershipMonth(membershipType.MembershipMonth).
		Save(context.Background())
	if err != nil {
		log.Fatal("Got an error when creating membershiptype. ERROR: ", err)
	}
	return createdMembershipType, nil
}

func (r *MembershipTypeRepositoryImpl) UpdateMembershipType(membershipType *ent.Membership_type) error {
	err := r.client.Membership_type.UpdateOneID(membershipType.ID).
		SetName(membershipType.Name).
		SetDescription(membershipType.Description).
		SetMembershipMonth(membershipType.MembershipMonth).
		SetPrice(membershipType.Price).
		Exec(context.Background())

	if err != nil {
		log.Fatal("Got an error when updating membership type. ERROR: ", err)
	}
	return nil
}

func (r *MembershipTypeRepositoryImpl) RemoveMembershipTypeByID(ID string) error {
	membershipTypeID, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	err = r.client.Membership_type.DeleteOneID(membershipTypeID).Exec(context.Background())
	return err
}
