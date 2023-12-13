package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/configs"
	"github.com/hmertakyatan/gymapi/handlers"
	"github.com/hmertakyatan/gymapi/middleware"
	"github.com/hmertakyatan/gymapi/repositories"
	"github.com/hmertakyatan/gymapi/services"
)

func main() {

	loadConfig, err := configs.LoadConfig("/home/server/api.env")
	if err != nil {
		log.Println("Could not load environments", err)
	}

	client := configs.ConnectionToDB(&loadConfig)

	defer client.Close()

	userRepository := repositories.UserRepositoryInit(client)

	customerRepository := repositories.CustomerRepositoryInit(client)

	customerService := services.CustomerServiceInit(customerRepository)

	customerHandler := handlers.CustomerHandlerInit(customerService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	customerAPI := router.Group("/api/customer")
	{
		customerAPI.GET("", customerHandler.HandleListCustomers)
		customerAPI.GET("/:id", customerHandler.HandleGetCustomerByID)
		customerAPI.POST("", customerHandler.HandleCreateCustomer)
		customerAPI.PUT("/:id", customerHandler.HandleUpdateCustomer)
		customerAPI.DELETE("/:id", customerHandler.HandleDeleteCustomerByID)
		customerAPI.GET("/debtors", customerHandler.HandleListDebtorCustomers)
	}

	server := &http.Server{
		Addr:    loadConfig.CustomerServerPort,
		Handler: router,
	}

	log.Printf("Customer microservice running on port: %s \n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Customer server failed to start:", err)
	}

}
