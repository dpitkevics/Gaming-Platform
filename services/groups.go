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

func GetGroupById(groupId int) *models.Group {
	var group models.Group

	db := database.GetDatabase()
	if notFound := db.First(&group, groupId).RecordNotFound(); notFound == true {
		return nil
	}

	return &group
}
