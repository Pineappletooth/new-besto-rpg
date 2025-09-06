package battle

import "github.com/google/uuid"

func processRound(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		for _, skill := range entity.ChosenSkills {
			battle.UseSkill(skill, entity)
		}
		entity.ChosenSkills = make([]string, 0)
	}
}

func New(entities []BattleEntity) *Battle {
	return &Battle{
		Id:       uuid.NewString(),
		entities: entities,
	}
}

func selectSkill(battle *Battle, battleId string, skills []string) {
	entity, ok := battle.getEntityById(battleId)
	if !ok {
		return
	}
	for _, skill := range skills {
		if _, ok := battle.GetSkill(skill); !ok {
			return
		}
	}
	entity.ChosenSkills = skills
	onSelectSkill(battle)
}

func onSelectSkill(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		if !entity.isDead() && len(entity.ChosenSkills) == 0 {
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
		teams[battle.entities[i].Team] = true
		if battle.entities[i].isDead() {
			deadTeams[battle.entities[i].Team] = true
		}
	}

	return len(teams)-len(deadTeams) <= 1
}

func end(battle *Battle) {
	//
}
