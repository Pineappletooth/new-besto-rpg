package battle

type Skill struct {
	name  string
	onUse func(battle *Battle, self *BattleEntity)
}

type Status struct {
	name     string
	priority int
	//effect   *Effect[Context]
}
