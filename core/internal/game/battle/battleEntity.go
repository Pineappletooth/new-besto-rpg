package battle

import (
	"pineappletooth/bestoRpg/internal/model"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id     string
	Stats  model.Stats
	Skills []string
}
type BattleEntity struct {
	Id           string
	Stats        model.Stats
	Status       []Status
	ChosenSkills []string
	Events       Events
	Team         int
	Base         BaseEntity
}

func NewFromBaseEntity(base BaseEntity) BattleEntity {
	return BattleEntity{
		Id:           base.Id,
		Stats:        base.Stats,
		Status:       make([]Status, 0),
		ChosenSkills: make([]string, 0),
		Events:       newEvents(),
		Base:         base,
		Team:         -1,
	}
}

func (b *BattleEntity) isDead() bool {
	return b.Stats.HP <= 0
}

func NewBattleEntityTest() BattleEntity {
	return NewFromBaseEntity(BaseEntity{
		Id: uuid.NewString(),
		Stats: model.Stats{
			HP:    10,
			Aggro: 1,
		},
		Skills: []string{"attack"},
	})
}

func NewBattleEntityFromCharacter(c model.Character) BattleEntity {
	stats := model.Stats{}
	skills := make([]string, 0)

	for _, item := range c.Equipment.Items {
		stats.AddStat(item.Stats)
		skills = append(skills, item.Skills...)
	}

	return NewFromBaseEntity(BaseEntity{
		Id:     c.Id,
		Stats:  stats,
		Skills: skills,
	})
}
