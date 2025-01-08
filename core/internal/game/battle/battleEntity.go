package battle



type BattleEntity struct {
	hp int
	skills map[string]Skill
	status []Status
}

func (s *BattleEntity) RollDiceT(dice DiceType) int{
	return s.RollDice(dices[dice])
}

func (s *BattleEntity) RollDice(dice []int) int{
	ctx := &onRollDiceContext{
		dice : dice,
	}
	for _, status := range s.status {
		if status.effect.onBeforeRollDice != nil {
			status.effect.onBeforeRollDice(ctx)
		}
		ctx.originalResult = RollDice(ctx.dice)
		ctx.result = ctx.originalResult
		if status.effect.onAfterRollDice != nil {
			status.effect.onAfterRollDice(ctx)
		}
	}
	return ctx.result
}

func (s *BattleEntity) Damage(damage int, target *BattleEntity){
	ctx := &onDamageContext{
		damage : damage,
		dealer : s,
		receiver: target,
	}
	for _, status := range s.status {
		if status.effect.onBeforeDamage != nil {
			status.effect.onBeforeDamage(ctx)
		}
		
		

		if status.effect.onAfterDamage != nil {
			status.effect.onAfterDamage(ctx)
		}
	}
}
