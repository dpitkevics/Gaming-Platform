package services

import (
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/dpitkevics/GamingPlatform/database"
)

func CreateGroupInvitation(user *models.User, invitedUser *models.User, group *models.Group) *models.GroupInvitation {
	groupInvitation := &models.GroupInvitation{
		OwnerId: user.ID,
		InvitedUserId: invitedUser.ID,
		GroupId: group.ID,
		Status: models.GROUP_INVITATION_STATUS_PENDING,

		Owner: user,
		InvitedUser: invitedUser,
		Group: group,
	}

	return groupInvitation
}

func GetGroupInvitationByIdAndInvitedUser(groupInvitationId int, invitedUser *models.User) *models.GroupInvitation {
	var groupInvitation models.GroupInvitation

	db := database.GetDatabase()
	if notFound := db.Where("id = ? AND invited_user_id = ?", groupInvitationId, invitedUser.ID).First(&groupInvitation).RecordNotFound(); notFound == true {
		return nil
	}

	return &groupInvitation
}

func AcceptGroupInvitation(groupInvitation *models.GroupInvitation) {
	AssignGroupToUser(groupInvitation.Group, groupInvitation.InvitedUser)

	groupInvitation.Status = models.GROUP_INVITATION_STATUS_CONFIRMED

	db := database.GetDatabase()
	db.Save(groupInvitation)
}

func DeclineGroupInvitation(groupInvitation *models.GroupInvitation) {
	groupInvitation.Status = models.GROUP_INVITATION_STATUS_DECLINED

	db := database.GetDatabase()
	db.Save(groupInvitation)
}
