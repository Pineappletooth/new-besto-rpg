package battle

import (
	"testing"
)

func TestFlow(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.dmg(dmgCtx{emitter: self, dmg: battle.rollDice(self, []int{3})})
		},
	}

	entity1.team = 1
	entity2.team = 2
	entity1.skills[name] = skill
	entity2.skills[name] = skill

	battle := New([]BattleEntity{entity1, entity2})

	selectSkill(battle, entity1.id, []string{name, name})
	selectSkill(battle, entity2.id, []string{name})

	if battle.entities[0].stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].stats.HP)
	}
	if battle.entities[1].stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].stats.HP)
	}
}

func TestSimulation(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.dmg(dmgCtx{emitter: self, dmg: battle.rollDice(self, []int{3})})
		},
	}

	entity1.team = 1
	entity2.team = 2
	entity1.skills[name] = skill
	entity2.skills[name] = skill
	entity1.chosenSkills = append(entity1.chosenSkills, name, name)
	entity2.chosenSkills = append(entity2.chosenSkills, name)

	battle := New([]BattleEntity{entity1, entity2})

	processRound(battle)
	if battle.entities[0].stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].stats.HP)
	}
	if battle.entities[1].stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].stats.HP)
	}
}

func TestSimulationWithEvents(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.dmg(dmgCtx{emitter: self, dmg: battle.rollDice(self, []int{3})})
		},
	}

	entity1.events.onBeforeRollDice.Subscribe(func(before onBeforeRollDiceContext, after onBeforeRollDiceContext) onBeforeRollDiceContext {
		after.Dice = []int{1}
		return after
	})

	entity2.events.onAfterRollDice.Subscribe(func(before onAfterRollDiceContext, after onAfterRollDiceContext) onAfterRollDiceContext {
		after.Result = before.Result + 1
		return after
	})

	entity1.team = 1
	entity2.team = 2
	entity1.skills[name] = skill
	entity2.skills[name] = skill
	entity1.chosenSkills = append(entity1.chosenSkills, name, name)
	entity2.chosenSkills = append(entity2.chosenSkills, name)

	battle := Battle{
		entities: []BattleEntity{entity1, entity2},
	}

	processRound(&battle)
	if battle.entities[0].stats.HP != 6 {
		t.Error("Expected entity 1 to have 6 HP, has", battle.entities[0].stats.HP)
	}
	if battle.entities[1].stats.HP != 8 {
		t.Error("Expected entity 2 to have 8 HP, has", battle.entities[1].stats.HP)
	}
}
