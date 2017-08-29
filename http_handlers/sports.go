package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/dpitkevics/GamingPlatform/database"
	"github.com/dpitkevics/GamingPlatform/models"
	"net/http"
	"github.com/dpitkevics/GamingPlatform/responses"
	"github.com/dpitkevics/GamingPlatform/services"
	"strconv"
)

func GetAllSportsAction(context *gin.Context) {
	db := database.GetDatabase()

	var sports []models.Sport
	db.Find(&sports)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: &sports,
	})
}

func GetSportAction(context *gin.Context) {
	sportId, err := strconv.Atoi(context.Param("sportId"))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	sport := services.GetSportById(sportId)
	if sport == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "Sport not found",
		})
		return
	}

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: sport,
	})
}
