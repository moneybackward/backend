package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/moneybackward/backend/utils/errors"
)

func GenerateToken(user_id uuid.UUID) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &errors.UnauthorizedError{Message: fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])}
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	splitBearers := strings.Split(bearerToken, " ")
	// "Bearer token"
	if len(splitBearers) == 2 {
		return splitBearers[1]
	}
	// "token"
	if len(splitBearers) == 1 {
		return splitBearers[0]
	}
	// idk lol
	return ""
}

func ExtractTokenID(c *gin.Context) (uuid.UUID, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &errors.UnauthorizedError{Message: fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])}
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		u_id, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return uuid.Nil, err
		}
		return u_id, nil
	}
	return uuid.Nil, nil
}
