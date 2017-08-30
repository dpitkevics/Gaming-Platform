package http_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dpitkevics/GamingPlatform/responses"
	"github.com/dpitkevics/GamingPlatform/services"
	"strconv"
)

func CreateGroupInvitationAction(context *gin.Context) {
	var request struct{
		GroupId int `json:"group_id"`
		UserId  int `json:"user_id"`
	}

	err := context.BindJSON(&request)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	group := services.GetGroupById(request.GroupId)
	if group == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "Group not found",
		})
		return
	}

	user := services.GetUserById(request.UserId)
	if user == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "User not found",
		})
		return
	}

	authenticatedUser := services.GetAuthenticatedUser(context)
	if user.ID == authenticatedUser.ID {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status: http.StatusBadRequest,
			Message: "Can't invite yourself to group",
		})
		return
	}

	groupInvitation := services.CreateGroupInvitation(authenticatedUser, user, group)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: groupInvitation,
	})
}

func GetGroupInvitationAction(context *gin.Context) {
	invitationId, err := strconv.Atoi(context.Param("invitationId"))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	groupInvitation := services.GetGroupInvitationByIdAndInvitedUser(invitationId, services.GetAuthenticatedUser(context))
	if groupInvitation == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invitation not found",
		})
		return
	}

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: groupInvitation,
	})
}

func AcceptGroupInvitationAction(context *gin.Context) {
	invitationId, err := strconv.Atoi(context.Param("invitationId"))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	groupInvitation := services.GetGroupInvitationByIdAndInvitedUser(invitationId, services.GetAuthenticatedUser(context))
	if groupInvitation == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invitation not found",
		})
		return
	}

	services.AcceptGroupInvitation(groupInvitation)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: groupInvitation,
	})
}

func DeclineGroupInvitationAction(context *gin.Context) {
	invitationId, err := strconv.Atoi(context.Param("invitationId"))
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	groupInvitation := services.GetGroupInvitationByIdAndInvitedUser(invitationId, services.GetAuthenticatedUser(context))
	if groupInvitation == nil {
		context.SecureJSON(http.StatusBadRequest, responses.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invitation not found",
		})
		return
	}

	services.DeclineGroupInvitation(groupInvitation)

	context.SecureJSON(http.StatusOK, responses.Response{
		Status: http.StatusOK,
		Data: groupInvitation,
	})
}
