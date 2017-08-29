package services

import (
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/dpitkevics/GamingPlatform/database"
)

func GetSportById(sportId int) *models.Sport {
	var sport models.Sport

	db := database.GetDatabase()
	db.First(&sport, sportId)

	return &sport
}
