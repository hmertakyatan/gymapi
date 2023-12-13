package services

import (
	"time"

	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/ent"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
	"github.com/hmertakyatan/gymapi/utils"
)

type UserService interface {
	GetListUsers() ([]*ent.User, error)
	AddNewUser(user requests.DtoUserRegisterRequest) (*ent.User, error)
	GetUserByID(ID string) (*ent.User, error)
	UpdateUserData(updatedUser requests.DtoUserUpdateRequest, existingUser *ent.User) error
	DeleteUserByID(ID string) error
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func UserServiceInit(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

// AddNewUser implements UserService.
func (us *UserServiceImpl) AddNewUser(user requests.DtoUserRegisterRequest) (*ent.User, error) {
	hashedPassword, err := utils.PasswordHasher(user.Password)
	helpers.ErrorPanic(err)

	newUser := &ent.User{
		Username:  user.Username,
		Password:  hashedPassword,
		Role:      user.Role,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	return us.userRepo.CreateUser(newUser)
}

// DeleteUserByID implements UserService.
func (us *UserServiceImpl) DeleteUserByID(ID string) error {
	return us.userRepo.RemoveUserByID(ID)
}

// GetListUsers implements UserService.
func (us *UserServiceImpl) GetListUsers() ([]*ent.User, error) {
	return us.userRepo.ReadAllUsers()
}

// GetUserByID implements UserService.
func (us *UserServiceImpl) GetUserByID(ID string) (*ent.User, error) {
	return us.userRepo.ReadUserByID(ID)
}

// UpdateUserData implements UserService.
func (us *UserServiceImpl) UpdateUserData(userUpdateRequest requests.DtoUserUpdateRequest, updatedUser *ent.User) error {
	updatedUser.Username = userUpdateRequest.Username
	updatedUser.Password = userUpdateRequest.Password
	updatedUser.Status = userUpdateRequest.Status
	updatedUser.UpdatedAt = time.Now().UTC()

	return us.userRepo.UpdateUser(updatedUser)
}
