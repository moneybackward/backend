package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type UserController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController() UserController {
	userService := services.NewUserService()

	return &userController{
		userService: userService,
	}
}

func (ctrl *userController) Add(ctx *gin.Context) {
	var input dto.UserDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.Create(&input)
	if err != nil {
		log.Panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (userCtrl *userController) List(ctx *gin.Context) {
	users, err := userCtrl.userService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
