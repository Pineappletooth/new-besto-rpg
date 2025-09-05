package model

type Skill struct {
	Name        string `json:"name" redis:"name"`
	Description string `json:"description" redis:"description"`
	Effect      string `json:"on_use" redis:"on_use"`
}
