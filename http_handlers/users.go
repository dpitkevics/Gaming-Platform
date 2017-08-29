package http_handlers

import (
	"github.com/dpitkevics/GamingPlatform/responses"
	"github.com/dpitkevics/GamingPlatform/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMeAction(context *gin.Context) {
	user := services.GetAuthenticatedUser(context)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data:   user,
	})
}
