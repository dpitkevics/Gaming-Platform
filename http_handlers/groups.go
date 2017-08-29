package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/dpitkevics/GamingPlatform/models"
	"net/http"
	"github.com/dpitkevics/GamingPlatform/responses"
	"github.com/dpitkevics/GamingPlatform/services"
)

func CreateGroupAction(context *gin.Context) {
	var group *models.Group = &models.Group{}

	err := context.BindJSON(group)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	group = services.CreateGroup(group)

	user := services.GetAuthenticatedUser(context)
	services.AssignGroupToUser(group, user)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: group,
	})
}
