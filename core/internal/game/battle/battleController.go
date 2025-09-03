package battle

func processRound(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		for j := range entity.chosenSkills {
			battle.entities[i].skills[battle.entities[i].chosenSkills[j]].onUse(battle, entity)
		}
		entity.chosenSkills = make([]string, 0)
	}
}

func New(battle *Battle) *Battle {
	return battle
}

func selectSkill(battle *Battle, entity *BattleEntity, skill []string) {
	entity.chosenSkills = skill
	onSelectSkill(battle)
}

func onSelectSkill(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		if !entity.isDead() && len(entity.chosenSkills) == 0 {
			return
		}
	}
	onRoundStart(battle)
}

func onRoundStart(battle *Battle) {
	processRound(battle)
	if checkAllDead(battle) {
		end(battle)
	}
	onRoundStart(battle)
}

func checkAllDead(battle *Battle) bool {
	for i := range battle.entities {
		if !battle.entities[i].isDead() {
			return false
		}
	}
	return true
}

func end(battle *Battle) {
	//
}
