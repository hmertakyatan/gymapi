package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/dto/responses"
	"github.com/hmertakyatan/gymapi/services"
	"github.com/hmertakyatan/gymapi/validators"
)

type MembershipHandler interface {
	HandleListMemberships(c *gin.Context)
	HandleGetMembershipByID(c *gin.Context)
	HandleCreateMembership(c *gin.Context)
	HandleUpdateMembership(c *gin.Context)
	HandleDeleteMembershipByID(c *gin.Context)
}
type MembershipHandlerImpl struct {
	MembershipService     services.MembershipService
	PaymentService        services.PaymentService
	CustomerService       services.CustomerService
	MembershipTypeService services.MembershipTypeService
}

func MembershipHandlerInit(membershipService services.MembershipService, paymentService services.PaymentService,
	customerService services.CustomerService, membershipTypeService services.MembershipTypeService) MembershipHandler {
	return &MembershipHandlerImpl{
		MembershipService:     membershipService,
		PaymentService:        paymentService,
		CustomerService:       customerService,
		MembershipTypeService: membershipTypeService,
	}
}

func (h *MembershipHandlerImpl) HandleListMemberships(c *gin.Context) {
	memberships, err := h.MembershipService.GetListMemberships()
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Müşteri listesi alınamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	c.JSON(http.StatusOK, memberships)
}

func (h *MembershipHandlerImpl) HandleGetMembershipByID(c *gin.Context) {
	membershipID := c.Param("id")
	membership, err := h.MembershipService.GetMembershipByID(membershipID)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Müşteri bulunamadı",
			Error:   err.Error(),
		}
		c.JSON(http.StatusNotFound, webResponse)
		return
	}
	c.JSON(http.StatusOK, membership)
}

func (h *MembershipHandlerImpl) HandleCreateMembership(c *gin.Context) {
	var newMembershipRequest requests.DtoMembershipRequest
	if err := c.ShouldBindJSON(&newMembershipRequest); err != nil {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Geçersiz istek formatı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	customer, err := h.CustomerService.GetCustomerByID(newMembershipRequest.CustomerID)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Müşteri bulunamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	membershipType, err := h.MembershipTypeService.GetMembershipTypeByID(newMembershipRequest.MembershipTypeID)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Üyelik tipi bulunumadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	if newMembershipRequest.AmountPaid > membershipType.Price {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Müşteriden, üyelik tipinin ücretinden fazla miktarda ödeme alınamaz.",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}
	if !validators.IsValidDate(newMembershipRequest.StartDateStr) {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Geçersiz tarih formatı. Lütfen gg/aa/yyyy formatında bir tarih giriniz.",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	addedMembership, err := h.MembershipService.AddNewMembership(newMembershipRequest)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Üyelik oluşturulamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	paymentDescription := fmt.Sprintf("%s - isimli musteriden, %s - uyelik tipi icin - %.2f TL - tutarinda ucret alindi", customer.FullName, membershipType.Name, addedMembership.AmountPaid)

	paymentRequest := requests.DtoPaymentRequest{
		Type:        "in",
		Description: paymentDescription,
		Amount:      addedMembership.AmountPaid,
	}

	paymentFromNewMembership, err := h.PaymentService.CreatePayment(paymentRequest)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Ödeme oluşturulamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := responses.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Üyelik ve ödeme işlemleri başarıyla gerçekleşti.",
		Data: map[string]interface{}{
			"membership": addedMembership,
			"payment":    paymentFromNewMembership,
		},
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *MembershipHandlerImpl) HandleUpdateMembership(c *gin.Context) {
	membershipID := c.Param("id")
	var updateMembershipRequest requests.DtoMembershipRequest
	if err := c.ShouldBindJSON(&updateMembershipRequest); err != nil {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Geçersiz istek formatı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}
	if !validators.IsValidDate(updateMembershipRequest.StartDateStr) {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Geçersiz tarih formatı. Lütfen gg/aa/yyyy formatında bir tarih giriniz.",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	existingMembership, err := h.MembershipService.GetMembershipByID(membershipID)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Güncellenecek olan üyelik bulunamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	if err := h.MembershipService.UpdateMembershipData(updateMembershipRequest, existingMembership); err != nil {
		webResponse := responses.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Üyeliği güncellerken bir hata oldu.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	webResponse := responses.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Üyelik başarıyla güncellendi.",
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *MembershipHandlerImpl) HandleDeleteMembershipByID(c *gin.Context) {
	membershipID := c.Param("id")
	err := h.MembershipService.DeleteMembershipByID(membershipID)
	if err != nil {
		webResponse := responses.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Silinecek olanacak müşteri bulunamadı.",
			Error:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	webResponse := responses.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Üyelik başarıyla silindi.",
	}

	c.JSON(http.StatusOK, webResponse)
}
