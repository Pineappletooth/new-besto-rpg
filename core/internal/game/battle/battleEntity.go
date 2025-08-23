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
}

func (s *BattleEntity) RollDiceT(dice DiceType) int {
	return s.RollDice(dices[dice])
}

func (s *BattleEntity) RollDice(dice []int) int {
	ctx := createEventContext(onRollDiceContext{
		dice:   dice,
		result: 0,
	})
	modifiedCtx := ctx.modified.(onRollDiceContext)

	print(modifiedCtx.result)

	for _, status := range s.status {
		if status.effect.onAfterEvent != nil {
			status.effect.onAfterEvent(&ctx)
		}
		ctx.originalResult = RollDice(ctx.dice)
		ctx.result = ctx.originalResult
		if status.effect.onAfterRollDice != nil {
			status.effect.onAfterRollDice(ctx)
		}
	}
	return ctx.result
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
