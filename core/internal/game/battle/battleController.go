package battle

func processRound(battle *Battle) {
	for i := range battle.entities {
		for j := range battle.entities[i].chosenSkills {
			battle.entities[i].skills[battle.entities[i].chosenSkills[j]].onUse(battle, &battle.entities[i])
		}
	}
}

func start(battle *Battle) {

}

func end(battle *Battle) {

}
