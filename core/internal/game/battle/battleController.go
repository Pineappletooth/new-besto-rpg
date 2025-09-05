package battle

import "github.com/google/uuid"

func processRound(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		for j := range entity.chosenSkills {
			battle.entities[i].skills[battle.entities[i].chosenSkills[j]].onUse(battle, entity)
		}
		entity.chosenSkills = make([]string, 0)
	}
}

func New(entities []BattleEntity) *Battle {
	return &Battle{
		id:       uuid.NewString(),
		entities: entities,
	}
}

func selectSkill(battle *Battle, string battleId, skill []string) {
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
	if checkEndBattle(battle) {
		end(battle)
	}
}

func checkEndBattle(battle *Battle) bool {
	teams := make(map[int]bool)
	deadTeams := make(map[int]bool)
	for i := range battle.entities {
		teams[battle.entities[i].team] = true
		if battle.entities[i].isDead() {
			deadTeams[battle.entities[i].team] = true
		}
	}

	return len(teams)-len(deadTeams) <= 1
}

func end(battle *Battle) {
	//
}
