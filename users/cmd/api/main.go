package main

import (
	"users/cmd/api/controller"
	"users/cmd/api/usecase"
	"users/internal/database"

	"github.com/forge-io/backend/lib/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil{
		panic(err)
	}
  
	userRepository := repository.NewUserRepository(dbConnection)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUseCase)

	server.GET("/users", userController.GetAllUsers)
	server.POST("/users", userController.CreateUser)

	server.Run(":8080")
}
