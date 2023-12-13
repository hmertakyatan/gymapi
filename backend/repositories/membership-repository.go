package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/membership"
)

type MembershipRepository interface {
	ReadAllMemberships() ([]*ent.Membership, error)
	ReadMembershipByID(ID string) (*ent.Membership, error)
	CreateMembership(m *ent.Membership, c *ent.Customer) (*ent.Membership, error)
	UpdateMembership(m *ent.Membership) error
	RemoveMembershipByID(ID string) error
}

type MembershipRepositoryImpl struct {
	client *ent.Client
}

func MembershipRepositoryInit(client *ent.Client) MembershipRepository {
	return &MembershipRepositoryImpl{client: client}
}

func (r *MembershipRepositoryImpl) ReadAllMemberships() ([]*ent.Membership, error) {
	memberships, err := r.client.Membership.Query().WithCustomer().All(context.Background())
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

func (r *MembershipRepositoryImpl) ReadMembershipByID(ID string) (*ent.Membership, error) {
	membershipID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatal("Got an error when parsing membership id. ERROR: ", err)
		return nil, err
	}
	membership, err := r.client.Membership.Query().Where(membership.ID(membershipID)).WithCustomer().First(context.Background())
	if err != nil {
		log.Fatal("Got an error when finding membership by id. ERROR: ", err)
		return nil, err
	}
	return membership, nil
}

func (r *MembershipRepositoryImpl) CreateMembership(m *ent.Membership, c *ent.Customer) (*ent.Membership, error) {
	createdMembership, err := r.client.Membership.Create().
		SetStartDate(m.StartDate).
		SetEndDate(m.EndDate).
		SetAmountPaid(m.AmountPaid).
		SetAmountRemaining(m.AmountRemaining).
		SetOperationDate(m.OperationDate).
		SetCustomer(c).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return createdMembership, nil
}

func (r *MembershipRepositoryImpl) UpdateMembership(changedMembership *ent.Membership) error {
	err := r.client.Membership.UpdateOneID(changedMembership.ID).
		SetStartDate(changedMembership.StartDate).
		SetEndDate(changedMembership.EndDate).
		SetAmountPaid(changedMembership.AmountPaid).
		SetAmountRemaining(changedMembership.AmountRemaining).
		SetOperationDate(changedMembership.OperationDate).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *MembershipRepositoryImpl) RemoveMembershipByID(ID string) error {
	membershipID, err := uuid.Parse(ID)
	if err != nil {
		return err
	}

	err = r.client.Membership.DeleteOneID(membershipID).Exec(context.Background())
	return err
}
