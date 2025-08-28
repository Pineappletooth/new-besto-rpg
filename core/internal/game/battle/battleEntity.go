package battle

import "github.com/google/uuid"

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
	rollDiceEvent Event[onRollDiceContext]
}

func NewBattleEntity() BattleEntity {
	return BattleEntity{
		rollDiceEvent: getRollDiceEvent(),
		id:            uuid.New(),
	}
}

func (s BattleEntity) RollDice(diceType DiceType) int {
	dice := GetDice(diceType)
	ctx := onRollDiceContext{
		dice:   dice,
		result: 0,
	}
	res := s.rollDiceEvent.emit(createEventContext(ctx))
	return res.result
}

func getRollDiceEvent() Event[onRollDiceContext] {
	event := NewEvent(onEventRollDice)

	return event
}

func onEventRollDice(ctx EventContext[onRollDiceContext]) EventContext[onRollDiceContext] {
	ctx.modified.result += RollDice(ctx.modified.dice)
	return ctx
}

func (s *BattleEntity) Damage(damage int, target *BattleEntity) {

}
