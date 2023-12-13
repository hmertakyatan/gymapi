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

	customerRepository := repositories.CustomerRepositoryInit(client)
	customerService := services.CustomerServiceInit(customerRepository)

	membershipTypeRepository := repositories.MembershipTypeRepositoryInit(client)
	membershipTypeService := services.MembershipTypeServiceInit(membershipTypeRepository)

	membershipRepository := repositories.MembershipRepositoryInit(client)
	membershipService := services.MembershipServiceInit(membershipTypeRepository, customerRepository, membershipRepository)

	membershipHandler := handlers.MembershipHandlerInit(membershipService, paymentService, customerService, membershipTypeService)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	membershipAPI := router.Group("/api/membership")
	{
		membershipAPI.GET("", membershipHandler.HandleListMemberships)
		membershipAPI.GET("/:id", membershipHandler.HandleGetMembershipByID)
		membershipAPI.POST("", membershipHandler.HandleCreateMembership)
		membershipAPI.PUT("/:id", membershipHandler.HandleUpdateMembership)
		membershipAPI.DELETE("/:id", membershipHandler.HandleDeleteMembershipByID)
	}

	server := &http.Server{
		Addr:    loadConfig.MembershipServerPort,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Membership server failed to start: ", err)
	}

	log.Printf("Membership microservice running on port: %s \n", server.Addr)
}
