package battle

type BattleContext struct {
	self    *BattleEntity
	allies  []*BattleEntity
	enemies []*BattleEntity
}

type Event string

const (
	battleStart = "battleStart"
	battleEnd   = "battleEnd"
	roundStart  = "roundStart"
	roundEnd    = "roundEnd"
	turnStart   = "turnStart"
	turnEnd     = "turnEnd"
	rollDice    = "rollDice"
	changeStat  = "changeStat"
)

type onRollDiceContext struct {
	dice   []int
	result int
}

func (on onRollDiceContext) getEvent() Event {
	return rollDice
}

type onChangeStatContext struct {
	stat     Stat
	change   int
	dealer   *BattleEntity
	receiver *BattleEntity
}

func (on *onChangeStatContext) getEvent() Event {
	return changeStat
}

func createEventContext[T Context](ctx T) EventContext[Context] {
	return EventContext[Context]{
		original: ctx,
		modified: ctx,
	}
}

type Context interface {
	getEvent() Event
}

type EventContext[T Context] struct {
	original T
	modified T
}

type Effect[T Context] struct {
	onBeforeEvent func(ctx *EventContext[T])
	onEvent       func(ctx *EventContext[T])
	onAfterEvent  func(ctx *EventContext[T])
}

type Skill struct {
	name  string
	onUse func(target *BattleEntity)
}

type Status struct {
	name     string
	priority int
	effect   *Effect[Context]
}
