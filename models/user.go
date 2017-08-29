package models

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"-"`

	Groups []*Group `json:"groups" gorm:"many2many:user_groups;"`
}
