package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/services"
)

type ProfitHandler interface {
	HandleProfitData(c *gin.Context)
}

type ProfitHandlerImpl struct {
	PaymentService services.PaymentService
}

func ProfitHandlerInit(paymentService services.PaymentService) ProfitHandler {
	return &ProfitHandlerImpl{
		PaymentService: paymentService,
	}
}

func (h *ProfitHandlerImpl) HandleProfitData(c *gin.Context) {
	incomePayments, err := h.PaymentService.ListIncomePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	expensePayments, err := h.PaymentService.ListExpensePayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalIncome := helpers.CalculateTotalAmount(incomePayments)
	totalExpense := helpers.CalculateTotalAmount(expensePayments)
	Profit := totalIncome - totalExpense
	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome, "total_expense": totalExpense, "Profit": Profit})
}
