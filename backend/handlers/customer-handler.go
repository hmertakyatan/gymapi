package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hmertakyatan/gymapi/dto/requests"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/services"
)

type CustomerHandler interface {
	HandleListCustomers(c *gin.Context)
	HandleGetCustomerByID(c *gin.Context)
	HandleCreateCustomer(c *gin.Context)
	HandleUpdateCustomer(c *gin.Context)
	HandleDeleteCustomerByID(c *gin.Context)
	HandleListDebtorCustomers(c *gin.Context)
}

type CustomerHandlerImpl struct {
	CustomerService services.CustomerService
}

func CustomerHandlerInit(customerService services.CustomerService) CustomerHandler {
	return &CustomerHandlerImpl{
		CustomerService: customerService,
	}
}

func (h *CustomerHandlerImpl) HandleListCustomers(c *gin.Context) {
	customers, err := h.CustomerService.GetListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandlerImpl) HandleGetCustomerByID(c *gin.Context) {
	customerID := c.Param("id")
	customer, err := h.CustomerService.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandlerImpl) HandleCreateCustomer(c *gin.Context) {
	var customerRequest requests.DtoCustomerRequest
	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if customerRequest.CustomerPictureURL != "" {
		filename := fmt.Sprintf("static/media/customer-images/%s.png", uuid.New())
		err := helpers.SaveBase64Image(customerRequest.CustomerPictureURL, filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Resim kaydedilemedi"})
			return
		}
		customerRequest.CustomerPictureURL = filename
	}

	customer, err := h.CustomerService.AddNewCustomer(customerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandlerImpl) HandleUpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")
	var updateCustomerRequest requests.DtoCustomerRequest
	if err := c.ShouldBindJSON(&updateCustomerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingCustomer, err := h.CustomerService.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if updateCustomerRequest.CustomerPictureURL != "" {
		filename := fmt.Sprintf("static/media/customer-images/%s.png", uuid.New())
		err := helpers.SaveBase64Image(updateCustomerRequest.CustomerPictureURL, filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Resim kaydedilemedi"})
			return
		}
		updateCustomerRequest.CustomerPictureURL = filename
	}

	err = h.CustomerService.UpdateCustomerData(updateCustomerRequest, existingCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Mesaj: ": "Musteri basariyle guncellendi."})
}

func (h *CustomerHandlerImpl) HandleDeleteCustomerByID(c *gin.Context) {
	customerID := c.Param("id")
	err := h.CustomerService.DeleteCustomerByID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *CustomerHandlerImpl) HandleListDebtorCustomers(c *gin.Context) {
	debtors, err := h.CustomerService.GetListDebtorCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, debtors)
}
