package battle

type Skill struct {
	name  string
	onUse func(target *BattleEntity)
}

type Status struct {
	name     string
	priority int
	//effect   *Effect[Context]
}
