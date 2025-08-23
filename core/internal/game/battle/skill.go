package battle

type BattleContext struct {
	self    *BattleEntity
	allies  []*BattleEntity
	enemies []*BattleEntity
}

type Event[T Context] struct {
	name        EventName
	subscribers []EventAction[T]
}

func NewEvent[T Context]() Event[T] {
	var ctx T
	return Event[T]{
		name:        ctx.getEventName(),
		subscribers: make([]EventAction[T], 0),
	}
}

type EventName string

const (
	battleStart EventName = "battleStart"
	battleEnd   EventName = "battleEnd"
	roundStart  EventName = "roundStart"
	roundEnd    EventName = "roundEnd"
	turnStart   EventName = "turnStart"
	turnEnd     EventName = "turnEnd"
	rollDice    EventName = "rollDice"
	changeStat  EventName = "changeStat"
)

type onRollDiceContext struct {
	dice   []int
	result int
}

func (ctx onRollDiceContext) getEventName() EventName {
	return rollDice
}

type onChangeStatContext struct {
	stat     Stat
	change   int
	dealer   *BattleEntity
	receiver *BattleEntity
}

func (on *onChangeStatContext) getEventName() EventName {
	return changeStat
}

func createEventContext[T Context](ctx T) EventContext[Context] {
	return EventContext[Context]{
		original: ctx,
		modified: ctx,
	}
}

func (event Event[T]) subscribe(action EventAction[T]) {
	event.subscribers = append(event.subscribers, action)
}

func (event Event[T]) emit(cxt *EventContext[T]) {
	for _, sub := range event.subscribers {
		sub.onAfterEvent(cxt)
	}
	for _, sub := range event.subscribers {
		sub.onEvent(cxt)
	}
	for _, sub := range event.subscribers {
		sub.onBeforeEvent(cxt)
	}
}

type Context interface {
	getEventName() EventName
}

type EventContext[T Context] struct {
	original T
	modified T
}

type EventAction[T Context] struct {
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
	//effect   *Effect[Context]
}
