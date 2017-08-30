package models

const GROUP_INVITATION_STATUS_PENDING = 1
const GROUP_INVITATION_STATUS_CONFIRMED = 2
const GROUP_INVITATION_STATUS_DECLINED = 3

type GroupInvitation struct {
	Model

	OwnerId       uint `json:"owner_id"`
	InvitedUserId uint `json:"invited_user_id"`
	GroupId       uint `json:"group_id"`
	Status        uint `json:"status"`

	Owner       *User `json:"owner"`
	InvitedUser *User `json:"invited_user"`
	Group       *Group `json:"group"`
}
