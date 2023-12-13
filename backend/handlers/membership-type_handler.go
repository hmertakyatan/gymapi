package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/services"
)

type MembershipTypeHandler interface {
	HandleListMembershipTypes(c *gin.Context)
	HandleGetMembershipTypeByID(c *gin.Context)
	HandleCreateMembershipType(c *gin.Context)
	HandleUpdateMembershipType(c *gin.Context)
	HandleDeleteMembershipTypeByID(c *gin.Context)
}

type MembershipTypeHandlerImpl struct {
	MembershipTypeService services.MembershipTypeService
}

func MembershipTypeHandlerInit(membershipTypeService services.MembershipTypeService) MembershipTypeHandler {
	return &MembershipTypeHandlerImpl{
		MembershipTypeService: membershipTypeService,
	}
}

func (h *MembershipTypeHandlerImpl) HandleListMembershipTypes(c *gin.Context) {
	membershipTypes, err := h.MembershipTypeService.GetListMembershipTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, membershipTypes)
}

func (h *MembershipTypeHandlerImpl) HandleGetMembershipTypeByID(c *gin.Context) {
	membershipTypeID := c.Param("id")
	membershipType, err := h.MembershipTypeService.GetMembershipTypeByID(membershipTypeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, membershipType)
}

func (h *MembershipTypeHandlerImpl) HandleCreateMembershipType(c *gin.Context) {
	var newMembershipTypeRequest requests.DtoMembershipTypeRequest
	if err := c.ShouldBindJSON(&newMembershipTypeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	membershipType, err := h.MembershipTypeService.AddNewMembershipType(newMembershipTypeRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, membershipType)
}

func (h *MembershipTypeHandlerImpl) HandleUpdateMembershipType(c *gin.Context) {
	membershipTypeID := c.Param("id")
	var updateMembershipTypeRequest requests.DtoMembershipTypeRequest
	if err := c.ShouldBindJSON(&updateMembershipTypeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingMembershipType, err := h.MembershipTypeService.GetMembershipTypeByID(membershipTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.MembershipTypeService.UpdateMembershipTypeData(updateMembershipTypeRequest, existingMembershipType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mesaj": "Uyelik tipi guncellendi"})
}

func (h *MembershipTypeHandlerImpl) HandleDeleteMembershipTypeByID(c *gin.Context) {
	membershipTypeID := c.Param("id")
	err := h.MembershipTypeService.DeleteMembershipTypeByID(membershipTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
