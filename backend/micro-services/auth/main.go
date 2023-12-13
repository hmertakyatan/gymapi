package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	authenticationService := services.AuthenticationServiceInit(userRepository, validate)
	authHandler := handlers.AuthenticationHandlerInit(authenticationService)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	authAPI := router.Group("/api/auth")
	{
		authAPI.POST("/login", authHandler.UserLogin)
		authAPI.GET("/refresh", authHandler.RefreshAccessToken)
		authAPI.GET("/logout", authHandler.LogoutUser)
	}

	server := &http.Server{
		Addr:    loadConfig.AuthServerPort,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Auth server failed to start:", err)
	}

	log.Printf("Auth microservice running on port: %s \n", server.Addr)

}
