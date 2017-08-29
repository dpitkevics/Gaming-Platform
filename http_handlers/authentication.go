package http_handlers

import (
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/dpitkevics/GamingPlatform/responses"
	"github.com/dpitkevics/GamingPlatform/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserAction(context *gin.Context) {
	var user *models.User = &models.User{}

	err := context.BindJSON(user)

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	services.CreateUser(user)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data:   user,
	})
}
