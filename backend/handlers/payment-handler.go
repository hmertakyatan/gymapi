package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/services"
)

type PaymentHandler interface {
	HandleListAllPayments(c *gin.Context)
	HandleGetPaymentByID(c *gin.Context)
	HandleCreatePayment(c *gin.Context)
	HandleUpdatePayment(c *gin.Context)
	HandleDeletePaymentByID(c *gin.Context)
	HandleListIncomePayments(c *gin.Context)
	HandleListExpensePayments(c *gin.Context)
}

type PaymentHandlerImpl struct {
	PaymentService services.PaymentService
}

func PaymentHandlerInit(paymentService services.PaymentService) PaymentHandler {
	return &PaymentHandlerImpl{
		PaymentService: paymentService,
	}
}

func (h *PaymentHandlerImpl) HandleListAllPayments(c *gin.Context) {
	payments, err := h.PaymentService.ListAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandlerImpl) HandleGetPaymentByID(c *gin.Context) {
	paymentID := c.Param("id")
	payment, err := h.PaymentService.GetPaymentByID(paymentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandlerImpl) HandleCreatePayment(c *gin.Context) {
	var paymentRequest requests.DtoPaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := h.PaymentService.CreatePayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

func (h *PaymentHandlerImpl) HandleUpdatePayment(c *gin.Context) {
	paymentID := c.Param("id")
	var paymentRequest requests.DtoPaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedPayment, err := h.PaymentService.UpdatePaymentByID(paymentID, paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPayment)
}

func (h *PaymentHandlerImpl) HandleDeletePaymentByID(c *gin.Context) {
	paymentID := c.Param("id")
	err := h.PaymentService.DeletePaymentByID(paymentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *PaymentHandlerImpl) HandleListIncomePayments(c *gin.Context) {
	incomePayments, err := h.PaymentService.ListIncomePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, incomePayments)
}

func (h *PaymentHandlerImpl) HandleListExpensePayments(c *gin.Context) {
	expensePayments, err := h.PaymentService.ListExpensePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, expensePayments)
}
