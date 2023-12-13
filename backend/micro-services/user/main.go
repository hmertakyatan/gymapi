package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/configs"
	"github.com/hmertakyatan/gymapi/handlers"
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

	userService := services.UserServiceInit(userRepository)

	userHandler := handlers.UserHandlerInit(userService)

	router := gin.Default()

	userAPI := router.Group("/api/user")
	{
		userAPI.GET("", userHandler.ListUsers)
		userAPI.GET("/:id", userHandler.GetUserByID)
		userAPI.POST("", userHandler.UserRegistration)
		userAPI.PUT("/:id", userHandler.UserUpdate)
		userAPI.DELETE("/:id", userHandler.DeleteUser)
	}

	server := &http.Server{
		Addr:    loadConfig.UserServerPort,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("User server failed to start:", err)
	}

	log.Printf("User microservice running on port: %s \n", server.Addr)
}
