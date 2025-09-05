package battle

import (
	"pineappletooth/bestoRpg/internal/model"

	"github.com/google/uuid"
)

type BattleEntity struct {
	Id            string
	stats         model.Stats
	originalStats model.Stats
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
		Id:     uuid.NewString(),
		events: newEvents(),
		stats: model.Stats{
			HP:    10,
			Aggro: 1,
		},
		originalStats: model.Stats{
			HP:    10,
			Aggro: 1,
		},
		skills: map[string]Skill{},
	}
}

func NewBattleEntityFromCharacter(c model.Character) BattleEntity {
	return BattleEntity{}
}
