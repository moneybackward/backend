package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type UserController interface {
	ListUsers(ctx *gin.Context)
	AddUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

// AddUser implements UserController.
func (ctrl *userController) AddUser(ctx *gin.Context) {
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

func (userCtrl *userController) ListUsers(ctx *gin.Context) {
	users, err := userCtrl.userService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func NewUserController() UserController {
	userService := services.NewUserService()

	return &userController{
		userService: userService,
	}
}
