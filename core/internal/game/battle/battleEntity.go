package battle

type Stat string

const (
	HP = "HP"
)

type Stats map[Stat]int

type BattleEntity struct {
	stats         Stats
	originalStats Stats
	skills        map[string]Skill
	status        []Status
	chosenSkills  []string
	rollDiceEvent Event[onRollDiceContext]
}

func getRollDiceEvent() Event[onRollDiceContext] {
	event := NewEvent[onRollDiceContext]()
	event.subscribe(EventAction[onRollDiceContext]{
		onEvent: onEventRollDice,
	})
	return event
}

func onEventRollDice(ctx *EventContext[onRollDiceContext]) {
	ctx.modified.result += RollDice(ctx.modified.dice)
}

func (s *BattleEntity) Damage(damage int, target *BattleEntity) {
	ctx := &onChangeStatContext{
		stat:     HP,
		change:   -damage,
		dealer:   s,
		receiver: target,
	}
	for _, status := range s.status {
		if status.effect.onAfterStatChange != nil {
			status.effect.onBeforeStatChange(ctx)
		}

		target.stats[HP] += ctx.change

		if status.effect.onAfterStatChange != nil {
			status.effect.onAfterStatChange(ctx)
		}
	}
}
