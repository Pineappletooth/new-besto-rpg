package battle

import "pineappletooth/bestoRpg/internal/game/event"

const BeforeRollDice event.EventName = "beforeRollDice"
const AfterRollDice event.EventName = "afterRollDice"

type onBeforeRollDiceContext struct {
	Dice []int
}

func (onBeforeRollDiceContext) GetEventName() event.EventName {
	return BeforeRollDice
}

type onAfterRollDiceContext struct {
	Dice   []int
	Result int
}

func (onAfterRollDiceContext) GetEventName() event.EventName {
	return AfterRollDice
}
