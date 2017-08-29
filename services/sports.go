package services

import (
	"github.com/dpitkevics/GamingPlatform/database"
	"github.com/dpitkevics/GamingPlatform/models"
)

func GetSportById(sportId int) *models.Sport {
	var sport models.Sport

	db := database.GetDatabase()
	if notFound := db.First(&sport, sportId).RecordNotFound(); notFound == true {
		return nil
	}

	return &sport
}
