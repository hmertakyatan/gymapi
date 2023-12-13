package main

import (
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

	paymentRepository := repositories.PaymentRepositoryInit(client)

	paymentService := services.PaymentServiceInit(paymentRepository)

	paymentHandler := handlers.PaymentHandlerInit(paymentService)

	profitHandler := handlers.ProfitHandlerInit(paymentService)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	paymentAPI := router.Group("api/payment")
	{
		paymentAPI.GET("", paymentHandler.HandleListAllPayments)
		paymentAPI.GET("/:id", paymentHandler.HandleGetPaymentByID)
		paymentAPI.POST("", paymentHandler.HandleCreatePayment)
		paymentAPI.PUT("/:id", paymentHandler.HandleUpdatePayment)
		paymentAPI.DELETE("/:id", paymentHandler.HandleDeletePaymentByID)
		paymentAPI.GET("/profit", profitHandler.HandleProfitData)
	}

	server := &http.Server{
		Addr:    loadConfig.PaymentServerPort,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Payment server failed to start:", err)
	}

	log.Printf("Payment microservice running on port: %s \n", server.Addr)

}
