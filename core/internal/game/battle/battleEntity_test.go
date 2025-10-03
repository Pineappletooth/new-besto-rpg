package battle

import (
	"testing"
)

func TestSimulationWithEvents(t *testing.T) {
	controller, battle := getTestBattle()

	attack := "attack"
	entity1 := battle.entities[0]
	entity2 := battle.entities[1]

	entity1.Events.OnBeforeRollDice.Subscribe(func(before OnBeforeRollDiceContext, after OnBeforeRollDiceContext) OnBeforeRollDiceContext {
		after.Dice = []int{1}
		return after
	})

	entity2.Events.OnAfterRollDice.Subscribe(func(before OnAfterRollDiceContext, after OnAfterRollDiceContext) OnAfterRollDiceContext {
		after.Result = before.Result + 1
		return after
	})

	entity1.ChosenSkills = append(entity1.ChosenSkills, attack, attack)
	entity2.ChosenSkills = append(entity2.ChosenSkills, attack)

	controller.processRound(battle)

	if battle.entities[0].Stats.HP != 6 {
		t.Error("Expected entity 1 to have 6 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 8 {
		t.Error("Expected entity 2 to have 8 HP, has", battle.entities[1].Stats.HP)
	}
}
