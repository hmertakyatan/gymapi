package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/ent/user"
)

type UserRepository interface {
	ReadAllUsers() ([]*ent.User, error)
	ReadUserByID(ID string) (*ent.User, error)
	ReadUserByUsername(username string) (*ent.User, error)
	CreateUser(user *ent.User) (*ent.User, error)
	UpdateUser(user *ent.User) error
	RemoveUserByID(ID string) error
}

func UserRepositoryInit(client *ent.Client) UserRepository {
	return &UserRepositoryImpl{client: client}
}

type UserRepositoryImpl struct {
	client *ent.Client
}

// ReadUserByUsername implements UserRepository.
func (r *UserRepositoryImpl) ReadUserByUsername(username string) (*ent.User, error) {
	user, err := r.client.User.Query().Where(user.Username(username)).First(context.Background())
	if err != nil {
		fmt.Println("Got an error when reading user by username. ERROR: ", err)
	}
	return user, nil
}

// CreateUser implements UserRepository.
func (r *UserRepositoryImpl) CreateUser(user *ent.User) (*ent.User, error) {
	createdUser, err := r.client.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetRole(user.Role).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		SetStatus(user.Status).
		Save(context.Background())

	if err != nil {
		fmt.Println("Got an error when creating user. ERROR: ", err)
	}

	return createdUser, nil
}

// ReadAllUsers implements UserRepository.
func (r *UserRepositoryImpl) ReadAllUsers() ([]*ent.User, error) {
	Users, err := r.client.User.Query().All(context.Background())
	if err != nil {
		fmt.Println("Got an error when reading all users. ERROR: ", err)
	}
	return Users, nil
}

// ReadUserByID implements UserRepository.
func (r *UserRepositoryImpl) ReadUserByID(ID string) (*ent.User, error) {
	userID, err := uuid.Parse(ID)
	if err != nil {
		fmt.Println("Got an error when parsing user id. ERROR: ", err)
	}

	user, err := r.client.User.Query().Where(user.ID(userID)).First(context.Background())

	if err != nil {
		fmt.Println("Got an error when reading user by user id. ERROR: ", err)
	}

	return user, nil
}

// RemoveUserByID implements UserRepository.
func (r *UserRepositoryImpl) RemoveUserByID(ID string) error {
	userID, err := uuid.Parse(ID)
	if err != nil {
		fmt.Println("Got an error when parsing user id. ERROR: ", err)
	}

	err = r.client.User.DeleteOneID(userID).Exec(context.Background())

	return err
}

// UpdateUser implements UserRepository.
func (r *UserRepositoryImpl) UpdateUser(user *ent.User) error {
	err := r.client.User.UpdateOneID(user.ID).
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetRole(user.Role).
		SetStatus(user.Status).
		SetUpdatedAt(time.Now().UTC()).
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}
