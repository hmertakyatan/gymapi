package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/services"
)

type PersonalTrainingHandler interface {
	HandlePersonalTrainingRegistration(c *gin.Context)
}

type PersonalTrainingHandlerImpl struct {
	CustomerService       services.CustomerService
	PaymentService        services.PaymentService
	PersonnelService      services.PersonnelService
	PersontrainingService services.PersonalTrainingService
}

func PersonalTrainingHandlerInit(customerService services.CustomerService, paymentService services.PaymentService,
	personnelService services.PersonnelService, personalTrainingService services.PersonalTrainingService) PersonalTrainingHandler {
	return &PersonalTrainingHandlerImpl{
		CustomerService:       customerService,
		PaymentService:        paymentService,
		PersonnelService:      personnelService,
		PersontrainingService: personalTrainingService,
	}
}

func (h *PersonalTrainingHandlerImpl) HandlePersonalTrainingRegistration(c *gin.Context) {

	var personalTrainingRegistrationRequest requests.DtoPersonalTrainingRegistartionRequest
	if err := c.ShouldBindJSON(&personalTrainingRegistrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.CustomerService.GetCustomerByID(personalTrainingRegistrationRequest.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	personnel, err := h.PersonnelService.GetPersonnelByID(personalTrainingRegistrationRequest.PersonnelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	personelRelCustomer, err := h.PersontrainingService.AddNewPersonalTrainingRegistration(personnel, customer, personalTrainingRegistrationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	paymentDescription := fmt.Sprintf("%s - isimli trainer ile %s isimli müşterinin personal tarining hizmeti için - %2.f TL ucret alindi", personnel.Name, customer.FullName, personelRelCustomer.Price)
	paymentRequest := requests.DtoPaymentRequest{
		Type:        "in",
		Description: paymentDescription,
		Amount:      personalTrainingRegistrationRequest.Price,
	}

	paymentFromNewPTRegistration, err := h.PaymentService.CreatePayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"personelRelCustomer": personelRelCustomer,
		"payment":             paymentFromNewPTRegistration,
	})

}
