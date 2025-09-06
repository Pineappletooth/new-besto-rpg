package battle

import (
	"pineappletooth/bestoRpg/internal/model"

	"github.com/google/uuid"
)

type BattleEntity struct {
	Id            string
	Stats         model.Stats
	OriginalStats model.Stats
	Skills        []string
	Status        []Status
	ChosenSkills  []string
	Events        Events
	Team          int
}

func (b *BattleEntity) isDead() bool {
	return b.Stats.HP <= 0
}

func NewBattleEntity() BattleEntity {
	return BattleEntity{
		Id:     uuid.NewString(),
		Events: newEvents(),
		Stats: model.Stats{
			HP:    10,
			Aggro: 1,
		},
		OriginalStats: model.Stats{
			HP:    10,
			Aggro: 1,
		},
		Skills: map[string]Skill{},
	}
}

func NewBattleEntityFromCharacter(c model.Character) BattleEntity {
	stats := model.Stats{}
	skills
	return BattleEntity{
		Id:     c.Id,
		Events: newEvents(),
		Stats:  c.Equipment.Items,
	}
}
