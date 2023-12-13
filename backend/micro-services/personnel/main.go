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

	personnelRepository := repositories.PersonnelRepositoryInit(client)

	personnelService := services.PersonnelServiceInit(personnelRepository)

	personnelHandler := handlers.PersonnelHandlerInit(personnelService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.DeserializeUser(userRepository))

	personnelAPI := router.Group("/api/personnel")
	{
		personnelAPI.GET("", personnelHandler.HandleListPersonnels)
		personnelAPI.GET("/:id", personnelHandler.HandleGetPersonnelByID)
		personnelAPI.POST("", personnelHandler.HandleCreatePersonnel)
		personnelAPI.PUT("/:id", personnelHandler.HandleUpdatePersonnel)
		personnelAPI.DELETE("/:id", personnelHandler.HandleDeletePersonnelByID)
	}

	server := &http.Server{
		Addr:    loadConfig.PersonnelServerPort,
		Handler: router,
	}

	log.Printf("Personnel microservice running on port: %s \n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Personnel server failed to start:", err)
	}

}
