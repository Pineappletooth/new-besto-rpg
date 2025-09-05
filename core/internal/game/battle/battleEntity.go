package battle

import (
	"github.com/google/uuid"
)

type Stat string

const (
	HP = "HP"
)

type Stats struct {
	HP    int
	Aggro int
}

type BattleEntity struct {
	id            string
	stats         Stats
	originalStats Stats
	skills        map[string]Skill
	status        []Status
	chosenSkills  []string
	events        Events
	team          int
}

func (b *BattleEntity) isDead() bool {
	return b.stats.HP <= 0
}

func NewBattleEntity() BattleEntity {
	return BattleEntity{
		id:     uuid.NewString(),
		events: newEvents(),
		stats: Stats{
			10,
			1,
		},
		originalStats: Stats{
			10,
			1,
		},
		skills: map[string]Skill{},
	}
}
