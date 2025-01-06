package battle

type Status string

type BattleEntity struct {
	hp int
	skills []Skill
	status []Status
}