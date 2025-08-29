package event

type Event[T Context] struct {
	name        EventName
	subscribers []EventAction[T]
	action      func(before T, after T) T
}

func NewEvent[T Context](action func(T, T) T) Event[T] {
	var ctx T
	return Event[T]{
		name:        ctx.GetEventName(),
		subscribers: make([]EventAction[T], 0),
		action:      action,
	}
}

type EventName string

func (event *Event[T]) Subscribe(action EventAction[T]) {
	event.subscribers = append(event.subscribers, action)
}

func (event *Event[T]) Emit(cxt T) T {
	before := cxt
	after := cxt
	for _, sub := range event.subscribers {
		after = sub.OnBefore(before, after)
	}
	before = event.action(before, after)
	after = before
	for _, sub := range event.subscribers {
		after = sub.OnAfter(before, after)
	}
	return after
}

type Context interface {
	GetEventName() EventName
}

type EventAction[T Context] struct {
	OnBefore func(before T, after T) T
	OnAfter  func(before T, after T) T
}
