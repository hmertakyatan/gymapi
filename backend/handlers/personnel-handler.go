package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/services"
)

type PersonnelHandler interface {
	HandleListPersonnels(c *gin.Context)
	HandleGetPersonnelByID(c *gin.Context)
	HandleCreatePersonnel(c *gin.Context)
	HandleUpdatePersonnel(c *gin.Context)
	HandleDeletePersonnelByID(c *gin.Context)
}

type PersonnelHandlerImpl struct {
	PersonnelService services.PersonnelService
}

func PersonnelHandlerInit(personnelService services.PersonnelService) PersonnelHandler {
	return &PersonnelHandlerImpl{
		PersonnelService: personnelService,
	}
}

func (h *PersonnelHandlerImpl) HandleListPersonnels(c *gin.Context) {
	personnels, err := h.PersonnelService.GetListPersonnels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personnels)
}

func (h *PersonnelHandlerImpl) HandleGetPersonnelByID(c *gin.Context) {
	personnelID := c.Param("id")
	personnel, err := h.PersonnelService.GetPersonnelByID(personnelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personnel)
}

func (h *PersonnelHandlerImpl) HandleCreatePersonnel(c *gin.Context) {
	var createPersonnelRequest requests.DtoPersonnelRequest
	if err := c.ShouldBindJSON(&createPersonnelRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personnel, err := h.PersonnelService.AddNewPersonnel(createPersonnelRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, personnel)
}

func (h *PersonnelHandlerImpl) HandleUpdatePersonnel(c *gin.Context) {
	personnelID := c.Param("id")
	var personnelUpdateRequest requests.DtoPersonnelRequest
	if err := c.ShouldBindJSON(&personnelUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPersonnel, err := h.PersonnelService.UpdatePersonnelData(personnelID, personnelUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPersonnel)
}

func (h *PersonnelHandlerImpl) HandleDeletePersonnelByID(c *gin.Context) {
	personnelID := c.Param("id")
	err := h.PersonnelService.DeletePersonnelByID(personnelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
