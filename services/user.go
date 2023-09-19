package services

import "github.com/moneybackward/backend/models"

func CreateUser(input models.CreateUserInput) models.User {
	// TODO: hash password
	user := models.User{Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	return user
}
