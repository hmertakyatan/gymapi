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

	paymentRepository := repositories.PaymentRepositoryInit(client)
	paymentService := services.PaymentServiceInit(paymentRepository)

	personnelRepository := repositories.PersonnelRepositoryInit(client)
	personnelService := services.PersonnelServiceInit(personnelRepository)

	ptRepository := repositories.PersonalTrainingRepositoryInit(client)

	ptService := services.PersonalTrainingServiceInit(ptRepository)

	ptHandler := handlers.PersonalTrainingHandlerInit(customerService, paymentService, personnelService, ptService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	ptAPI := router.Group("/api/pt")
	{
		ptAPI.POST("", ptHandler.HandlePersonalTrainingRegistration)
	}

	server := &http.Server{
		Addr:    loadConfig.PersonalTrainingServerPort,
		Handler: router,
	}

	log.Printf("Personnel microservice running on port: %s \n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Personnel server failed to start:", err)
	}

}
