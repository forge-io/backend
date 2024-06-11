package controller

import (
	"net/http"
	"users/cmd/api/usecase"

	"github.com/forge-io/backend/lib/models"
	"github.com/gin-gonic/gin"
)

type userController struct{
		userUseCase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController{
	return userController{
		userUseCase: usecase,
	}
}

func (u *userController) GetAllUsers(ctx *gin.Context){

	users, err := u.userUseCase.GetAllUsers()
	if(err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
	}
	
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) CreateUser(ctx *gin.Context){
	var user models.User
	err := ctx.BindJSON(&user)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUseCase.CreateUser(user)
	
	if err!= nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)
}