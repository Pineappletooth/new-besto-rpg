package battle

type BattleContext struct {
	self *BattleEntity
	allies []*BattleEntity
	enemies []*BattleEntity
}

type RollDiceContext struct {
	
}

type BattleArg struct {

}

type Skill struct {
	onBattleStart func()
	onBeforeRound func()
	onUse func(target BattleEntity)
	onBeforeRollDice func()
	onAfterRollDice func()
	onBeforeDamage func()
	onAfterDamage func()
	onAfterRound func()
}


