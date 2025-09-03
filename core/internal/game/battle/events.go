package battle

import "pineappletooth/bestoRpg/internal/game/event"

const BeforeRollDice event.EventName = "beforeRollDice"
const AfterRollDice event.EventName = "afterRollDice"

type Events struct {
	onBeforeRollDice event.Event[onBeforeRollDiceContext]
	onAfterRollDice  event.Event[onAfterRollDiceContext]
}

func newEvents() Events {
	return Events{
		onBeforeRollDice: event.New[onBeforeRollDiceContext](),
		onAfterRollDice:  event.New[onAfterRollDiceContext](),
	}
}

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
