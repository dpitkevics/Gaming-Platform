package services

import (
	"github.com/appleboy/gin-jwt"
	"github.com/dpitkevics/GamingPlatform/database"
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetUserById(userId int) *models.User {
	var user models.User

	db := database.GetDatabase()
	if notFound := db.First(&user, userId).RecordNotFound(); notFound == true {
		return nil
	}

	return &user
}

func GetUserByUsername(username string) *models.User {
	var user models.User

	db := database.GetDatabase()
	if notFound := db.Where("username = ?", username).First(&user).RecordNotFound(); notFound == true {
		return nil
	}

	return &user
}

func GetUserByUsernameAndPassword(username string, password string) *models.User {
	user := GetUserByUsername(username)
	if user == nil {
		return nil
	}

	if CheckPasswordHash(password, user.Password) {
		return user
	}

	return nil
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

	user.Password = password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	db := database.GetDatabase()
	db.Create(user)

	return user, nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashedPassword(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(newPassword), nil
}
