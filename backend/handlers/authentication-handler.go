package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/configs"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/dto/responses"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/services"
	"github.com/hmertakyatan/gymapi/utils"
)

type AuthenticationHandler interface {
	UserLogin(c *gin.Context)
	RefreshAccessToken(c *gin.Context)
	LogoutUser(c *gin.Context)
}

type AuthenticationHandlerImpl struct {
	AuthService services.AuthenticationService
}

func AuthenticationHandlerInit(authService services.AuthenticationService) AuthenticationHandler {
	return &AuthenticationHandlerImpl{
		AuthService: authService,
	}
}

// UserLogin implements AuthenticationHandler.
func (h *AuthenticationHandlerImpl) UserLogin(c *gin.Context) {
	userLoginRequest := requests.DtoUserLoginRequest{}
	err := c.ShouldBindJSON(&userLoginRequest)
	helpers.ErrorPanic(err)

	access_token, refresh_token, err := h.AuthService.Login(userLoginRequest)

	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := responses.UserLoginResponse{
		TokenType:   "Bearer",
		AccessToken: access_token,
	}

	webResponse := responses.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Kullanici girisi basarili",
		Data:    resp,
	}
	config, err := configs.LoadConfig("/home/server/api.env")

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, webResponse)
}

// RefreshAccessToken implements AuthenticationHandler.
func (h *AuthenticationHandlerImpl) RefreshAccessToken(c *gin.Context) {
	refreshtoken, err := c.Cookie("refresh_token")
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusForbidden,
			Status:  "Forbidden",
			Message: "Could not refresh token",
		}
		c.AbortWithStatusJSON(http.StatusForbidden, webResponse)
		return
	}

	config, err := configs.LoadConfig("/home/server/api.env")
	if err != nil {
		log.Println("Could not load environments", err)
	}

	sub, err := utils.ValidateToken(refreshtoken, config.RefreshTokenSecret)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusForbidden,
			Status:  "Forbidden",
			Message: "Refresh token is not valid.",
		}
		c.AbortWithStatusJSON(http.StatusForbidden, webResponse)
		return
	}

	userID := sub.(string)
	helpers.ErrorPanic(err)

	access_token, err := utils.GenerateToken(config.AccessTokenExpiredIn, userID, config.AccessTokenSecret)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusForbidden,
			Status:  "Forbidden",
			Message: "Could not generate access token.",
		}
		c.AbortWithStatusJSON(http.StatusForbidden, webResponse)
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

}

func (h *AuthenticationHandlerImpl) LogoutUser(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
