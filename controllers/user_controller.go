package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/models/dto"
	"github.com/moneybackward/backend/services"
	"github.com/rs/zerolog/log"
)

type UserController interface {
	List(*gin.Context)
	Register(*gin.Context)
	Login(*gin.Context)
	Delete(*gin.Context)
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

// @Summary Register a user
// @Tags auth
// @Accept json
// @Param category body dto.UserRegisterDTO true "User"
// @Success 201 {object} models.User
// @Router /auth/register [post]
func (ctrl *userController) Register(ctx *gin.Context) {
	var input dto.UserRegisterDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.userService.Create(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// @Summary Login a user
// @Tags auth
// @Accept json
// @Param category body dto.UserLoginDTO true "User"
// @Success 200 {object} nil
// @Router /auth/login [post]
func (userCtrl *userController) Login(ctx *gin.Context) {
	var input dto.UserLoginDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userCtrl.userService.Login(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": token})
}

// @Summary List users
// @Tags users
// @Success 200 {object} []models.User
// @Router /users [get]
// @Security BearerAuth
func (userCtrl *userController) List(ctx *gin.Context) {
	users, err := userCtrl.userService.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("Error finding users")
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
