package events

import (
	"pineappletooth/bestoRpg/internal/game/event"
	"pineappletooth/bestoRpg/internal/game/utils"
)

const RollDice event.EventName = "rollDice"

func GetRollDiceEvent() event.Event[OnRollDiceContext] {
	return event.NewEvent(onEventRollDice)
}

func onEventRollDice(before OnRollDiceContext, after OnRollDiceContext) OnRollDiceContext {
	after.Result += utils.RollDice(after.Dice)
	return after
}

type OnRollDiceContext struct {
	Dice   []int
	Result int
}

func (ctx OnRollDiceContext) GetEventName() event.EventName {
	return RollDice
}
