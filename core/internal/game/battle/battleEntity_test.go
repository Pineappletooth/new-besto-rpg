package battle

import (
	"testing"
)

func TestSimulationWithEvents(t *testing.T) {
	controller, battle := getTestBattle()

	name := "attack"
	entity1 := battle.entities[0]
	entity2 := battle.entities[1]

	entity1.Events.onBeforeRollDice.Subscribe(func(before onBeforeRollDiceContext, after onBeforeRollDiceContext) onBeforeRollDiceContext {
		after.Dice = []int{1}
		return after
	})

	entity2.Events.onAfterRollDice.Subscribe(func(before onAfterRollDiceContext, after onAfterRollDiceContext) onAfterRollDiceContext {
		after.Result = before.Result + 1
		return after
	})

	entity1.ChosenSkills = append(entity1.ChosenSkills, name, name)
	entity2.ChosenSkills = append(entity2.ChosenSkills, name)

	controller.processRound(battle)

	if battle.entities[0].Stats.HP != 6 {
		t.Error("Expected entity 1 to have 6 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 8 {
		t.Error("Expected entity 2 to have 8 HP, has", battle.entities[1].Stats.HP)
	}
}
