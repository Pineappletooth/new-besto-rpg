package battle

import (
	"testing"
)

func TestSimulation(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.dmg(self, battle.rollDice(self, []int{3}))
		},
	}

	entity1.skills[name] = skill
	entity2.skills[name] = skill
	entity1.chosenSkills = append(entity1.chosenSkills, name, name)
	entity2.chosenSkills = append(entity2.chosenSkills, name, name)

	battle := Battle{
		entities: []BattleEntity{entity1, entity2},
	}

	battle.processRound()
	if battle.entities[0].stats.HP != 4 {
		t.Error("Expected entity 1 to have 4 HP")
	}
	if battle.entities[1].stats.HP != 7 {
		t.Error("Expected entity 2 to have 7 HP")
	}
}
