package services

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/hmertakyatan/gymapi/configs"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
	"github.com/hmertakyatan/gymapi/utils"
)

type AuthenticationService interface {
	Login(user requests.DtoUserLoginRequest) (string, string, error)
}

type AuthenticationServiceImpl struct {
	userRepo repositories.UserRepository
	Validate *validator.Validate
}

func AuthenticationServiceInit(userRepo repositories.UserRepository, Validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		userRepo: userRepo,
		Validate: Validate,
	}
}

// Login implements AuthenticationService.
func (as *AuthenticationServiceImpl) Login(user requests.DtoUserLoginRequest) (string, string, error) {
	newUser, err := as.userRepo.ReadUserByUsername(user.Username)
	if err != nil {
		log.Println("Invalid username")
		return "", "", err
	}

	config, err := configs.LoadConfig("/home/server/api.env")
	if err != nil {
		log.Println("Could not loading configs.")
		return "", "", err
	}

	verifyErr := utils.VerifyPassword(newUser.Password, user.Password)
	if verifyErr != nil {
		log.Println("Invalid password")
		return "", "", verifyErr
	}

	access_token, err := utils.GenerateToken(config.AccessTokenExpiredIn, newUser.ID, config.AccessTokenSecret)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := utils.GenerateToken(config.RefreshTokenExpiredIn, newUser.ID, config.RefreshTokenSecret)
	if err != nil {
		return "", "", err
	}
	helpers.ErrorPanic(err)

	return access_token, refresh_token, nil

}
