package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/services"
)

type UserHandler interface {
	ListUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	UserRegistration(c *gin.Context)
	UserUpdate(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func UserHandlerInit(userService services.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

// DeleteUser implements UserHandler.
func (h *UserHandlerImpl) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	err := h.UserService.DeleteUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetUserByID implements UserHandler.
func (h *UserHandlerImpl) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ListUsers implements UserHandler.
func (h *UserHandlerImpl) ListUsers(c *gin.Context) {
	users, err := h.UserService.GetListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UserRegistration implements UserHandler.
func (h *UserHandlerImpl) UserRegistration(c *gin.Context) {
	userRegisterRequest := requests.DtoUserRegisterRequest{}
	if err := c.ShouldBindJSON(&userRegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.AddNewUser(userRegisterRequest)
	helpers.ErrorPanic(err)

	c.JSON(http.StatusOK, user)

}

// UserUpdate implements UserHandler.
func (h *UserHandlerImpl) UserUpdate(c *gin.Context) {
	userID := c.Param("id")
	userUpdateRequest := requests.DtoUserUpdateRequest{}
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	existingUser, err := h.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = h.UserService.UpdateUserData(userUpdateRequest, existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Mesaj: ": "Kullanici basariyle guncellendi."})
}
