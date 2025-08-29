package battle

import (
	"pineappletooth/bestoRpg/internal/game/battle/events"
	"pineappletooth/bestoRpg/internal/game/event"
	"pineappletooth/bestoRpg/internal/game/utils"

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
	rollDiceEvent event.Event[events.OnRollDiceContext]
}

func NewBattleEntity() BattleEntity {
	return BattleEntity{
		rollDiceEvent: events.GetRollDiceEvent(),
		id:            uuid.New(),
	}
}

func (s *BattleEntity) RollDice(diceType utils.DiceType) int {
	dice := utils.GetDice(diceType)
	ctx := events.OnRollDiceContext{
		Dice:   dice,
		Result: 0,
	}
	res := s.rollDiceEvent.Emit(ctx)
	return res.Result
}

func (s *BattleEntity) Damage(damage int, target *BattleEntity) {

}
