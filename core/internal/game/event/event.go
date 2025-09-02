package event

type Event[T Context] struct {
	name        EventName
	subscribers []Action[T]
}

func New[T Context]() Event[T] {
	var ctx T
	return Event[T]{
		name:        ctx.GetEventName(),
		subscribers: make([]Action[T], 0),
	}
}

type EventName string

func (event *Event[T]) Subscribe(action Action[T]) {
	event.subscribers = append(event.subscribers, action)
}

func (event *Event[T]) Emit(cxt T) T {
	before := cxt
	after := cxt
	for _, action := range event.subscribers {
		after = action(before, after)
	}
	return after
}

type Context interface {
	GetEventName() EventName
}

type Action[T Context] func(before T, after T) T
