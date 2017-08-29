package services

import (
	"github.com/appleboy/gin-jwt"
	"github.com/dpitkevics/GamingPlatform/database"
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetUserByUsername(username string) *models.User {
	var user models.User

	db := database.GetDatabase()
	db.Where("username = ?", username).First(&user)

	return &user
}

func GetUserByUsernameAndPassword(username string, password string) *models.User {
	user := GetUserByUsername(username)
	if user == nil {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil
	}

	return user
}

func GetAuthenticatedUser(context *gin.Context) *models.User {
	claims := jwt.ExtractClaims(context)
	return GetUserByUsername(claims["id"].(string))
}

func CreateUser(user *models.User) (*models.User, error) {
	password, err := GetHashedPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = string(password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	db := database.GetDatabase()
	db.Create(user)

	return user, nil
}

func GetHashedPassword(password string) (string, error) {
	defer clear(password)

	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(newPassword), nil
}

func clear(s string) {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}
