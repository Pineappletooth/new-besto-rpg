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
			battle.Dmg(dmgCtx{Emitter: self, Dmg: battle.RollDice(self, []int{3})})
		},
	}

	entity1.Team = 1
	entity2.Team = 2
	entity1.Skills[name] = skill
	entity2.Skills[name] = skill

	battle := New([]BattleEntity{entity1, entity2})

	selectSkill(battle, entity1.Id, []string{name, name})
	selectSkill(battle, entity2.Id, []string{name})

	if battle.entities[0].Stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].Stats.HP)
	}
}

func TestSimulation(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.Dmg(dmgCtx{Emitter: self, Dmg: battle.RollDice(self, []int{3})})
		},
	}

	entity1.Team = 1
	entity2.Team = 2
	entity1.Skills[name] = skill
	entity2.Skills[name] = skill
	entity1.ChosenSkills = append(entity1.ChosenSkills, name, name)
	entity2.ChosenSkills = append(entity2.ChosenSkills, name)

	battle := New([]BattleEntity{entity1, entity2})

	processRound(battle)
	if battle.entities[0].Stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].Stats.HP)
	}
}

func TestSimulationWithEvents(t *testing.T) {
	entity1 := NewBattleEntity()
	entity2 := NewBattleEntity()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.Dmg(dmgCtx{Emitter: self, Dmg: battle.RollDice(self, []int{3})})
		},
	}

	entity1.Events.onBeforeRollDice.Subscribe(func(before onBeforeRollDiceContext, after onBeforeRollDiceContext) onBeforeRollDiceContext {
		after.Dice = []int{1}
		return after
	})

	entity2.Events.onAfterRollDice.Subscribe(func(before onAfterRollDiceContext, after onAfterRollDiceContext) onAfterRollDiceContext {
		after.Result = before.Result + 1
		return after
	})

	entity1.Team = 1
	entity2.Team = 2
	entity1.Skills[name] = skill
	entity2.Skills[name] = skill
	entity1.ChosenSkills = append(entity1.ChosenSkills, name, name)
	entity2.ChosenSkills = append(entity2.ChosenSkills, name)

	battle := Battle{
		entities: []BattleEntity{entity1, entity2},
	}

	processRound(&battle)
	if battle.entities[0].Stats.HP != 6 {
		t.Error("Expected entity 1 to have 6 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 8 {
		t.Error("Expected entity 2 to have 8 HP, has", battle.entities[1].Stats.HP)
	}
}
