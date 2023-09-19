package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneybackward/backend/models"
	"github.com/moneybackward/backend/services"
)

func ListUsers(ctx *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(ctx *gin.Context) {
	var input models.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := services.CreateUser(input)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
