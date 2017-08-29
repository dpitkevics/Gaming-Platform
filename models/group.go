package models

type Group struct {
	Model

	Name string `json:"name" binding:"required,min=3"`
}
