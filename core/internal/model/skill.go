package model

type Skill struct {
	Name        string `json:"name" redis:"name"`
	Description string `json:"description" redis:"description"`
	Action      string `json:"action" redis:"action"`
}
