package battle

type BattleContext struct {
	self    *BattleEntity
	allies  []*BattleEntity
	enemies []*BattleEntity
}

type onRollDiceContext struct {
	dice           []int
	originalResult int
	result         int
}

type onChangeStatContext struct {
	stat     Stat
	change   int
	dealer   *BattleEntity
	receiver *BattleEntity
}

type Effect struct {
	onBattleStart      func()
	onBeforeRound      func()
	onBeforeTurn       func()
	onBeforeRollDice   func(ctx *onRollDiceContext)
	onAfterRollDice    func(ctx *onRollDiceContext)
	onBeforeStatChange func(ctx *onChangeStatContext)
	onAfterStatChange  func(ctx *onChangeStatContext)
	onAfterTurn        func()
	onAfterRound       func()
}

type Skill struct {
	name  string
	onUse func(target *BattleEntity)
}

type Status struct {
	name     string
	priority int
	effect   *Effect
}
