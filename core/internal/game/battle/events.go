package battle

import "pineappletooth/bestoRpg/internal/game/event"

const BeforeRollDice event.EventName = "beforeRollDice"
const AfterRollDice event.EventName = "afterRollDice"
const BeforeDmg event.EventName = "beforeDmg"

type Events struct {
	OnBeforeRollDice event.Event[OnBeforeRollDiceContext]
	OnAfterRollDice  event.Event[OnAfterRollDiceContext]
	OnBeforeDmg      event.Event[OnBeforeDmgContext]
}

func newEvents() Events {
	return Events{
		OnBeforeRollDice: event.New[OnBeforeRollDiceContext](),
		OnAfterRollDice:  event.New[OnAfterRollDiceContext](),
		OnBeforeDmg:      event.New[OnBeforeDmgContext](),
	}
}

type OnBeforeRollDiceContext struct {
	Dice []int
}

func (OnBeforeRollDiceContext) GetEventName() event.EventName {
	return BeforeRollDice
}

type OnAfterRollDiceContext struct {
	Dice   []int
	Result int
}

func (OnAfterRollDiceContext) GetEventName() event.EventName {
	return AfterRollDice
}

type OnBeforeDmgContext struct {
	Emitter *BattleEntity
	Target  *BattleEntity
	Dmg     int
}

func (OnBeforeDmgContext) GetEventName() event.EventName {
	return BeforeDmg
}
