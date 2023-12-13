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

	membershipTypeRepository := repositories.MembershipTypeRepositoryInit(client)

	membershipTypeService := services.MembershipTypeServiceInit(membershipTypeRepository)

	membershipTypeHandler := handlers.MembershipTypeHandlerInit(membershipTypeService)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	membershipTypeAPI := router.Group("/api/membership-type")
	{
		membershipTypeAPI.GET("", membershipTypeHandler.HandleListMembershipTypes)
		membershipTypeAPI.GET("/:id", membershipTypeHandler.HandleGetMembershipTypeByID)
		membershipTypeAPI.POST("", membershipTypeHandler.HandleCreateMembershipType)
		membershipTypeAPI.PUT("/:id", membershipTypeHandler.HandleUpdateMembershipType)
		membershipTypeAPI.DELETE("/:id", membershipTypeHandler.HandleDeleteMembershipTypeByID)
	}

	server := &http.Server{
		Addr:    loadConfig.MembershipTypeServerPort,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Membership Type server failed to start: ", err)
	}

	log.Printf("Membership Type microservice running on port: %s \n", server.Addr)

}
