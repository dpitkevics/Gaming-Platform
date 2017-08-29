package services

import (
	"github.com/dpitkevics/GamingPlatform/models"
	"time"
	"github.com/dpitkevics/GamingPlatform/database"
)

func CreateGroup(group *models.Group) *models.Group {
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	db := database.GetDatabase()
	db.Create(group)

	return group
}

func AssignGroupToUser(group *models.Group, user *models.User) {
	user.Groups = append(user.Groups, group)

	db := database.GetDatabase()
	db.Save(user)
}
