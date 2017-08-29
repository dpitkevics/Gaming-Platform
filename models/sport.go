package models

type Sport struct {
	Model

	Name              string `json:"name"`
	PlayerCountInTeam uint   `json:"player_count_in_team"`
}
