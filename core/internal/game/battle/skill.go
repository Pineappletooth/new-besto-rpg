package battle

type BattleContext struct {
	self *BattleEntity
	allies []*BattleEntity
	enemies []*BattleEntity
}

type onRollDiceContext struct {
	dice []int
	originalResult int
	result int
}


type onDamageContext struct {
	damage int
	dealer *BattleEntity
	receiver *BattleEntity
}


type Event struct {
	name string
}

type Context interface{
	onDamageContext | onRollDiceContext
}
type Effect2[T Context] struct {
	onBeforeEvent func(event T)
	onEvent func(event T)
	onAfterEvent func(event T)
}

func TriggerBeforeEvent() {
	a := Effect2[onDamageContext]{
		
	}
}
type Effect struct {
	onBattleStart func()
	onBeforeRound func()
	onBeforeTurn func()
	onBeforeRollDice func(ctx *onRollDiceContext)
	onAfterRollDice func(ctx *onRollDiceContext)
	onBeforeDamage func(ctx *onDamageContext)
	onAfterDamage func(ctx *onDamageContext)
	onAfterTurn func()
	onAfterRound func()
}

type Skill struct {
	name string
	onUse func(target *BattleEntity)
}

type Status struct {
	name string
	priority int
	effect *Effect
}