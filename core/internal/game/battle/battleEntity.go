package battle

import (
	"github.com/google/uuid"
)

type Stat string

const (
	HP = "HP"
)

type Stats map[Stat]int

type BattleEntity struct {
	id            uuid.UUID
	stats         Stats
	originalStats Stats
	skills        map[string]Skill
	status        []Status
	chosenSkills  []string
	events        Events
}

func NewBattleEntity() BattleEntity {
	return BattleEntity{
		id:     uuid.New(),
		events: newEvents(),
	}
}
