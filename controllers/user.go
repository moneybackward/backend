package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
)

type UserController interface {
	List(ctx *gin.Context)
	Add(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

// @Summary Add a user
// @Tags users
// @Accept json
// @Param category body dto.UserDTO true "User"
// @Success 201 {object} models.User
// @Router /users [post]
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

// @Summary List users
// @Tags users
// @Success 200 {object} []models.User
// @Router /users [get]
func (userCtrl *userController) List(ctx *gin.Context) {
	users, err := userCtrl.userService.FindAll()
	if err != nil {
		log.Panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary Delete a user
// @Tags users
// @Success 204 {object} nil
func (userCtrl *userController) Delete(ctx *gin.Context) {
	userId := uuid.MustParse(ctx.Param("id"))
	err := userCtrl.userService.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
