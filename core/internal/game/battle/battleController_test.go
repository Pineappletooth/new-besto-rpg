package battle

import (
	"pineappletooth/bestoRpg/internal/model"
	"testing"
)

type mockStatusPersistence struct {
}

func (m mockStatusPersistence) GetStatus(status string) (model.Status, error) {
	return model.Status{
		Name:        status,
		Description: status,
		Action: `skill.OnUse = function (battle, entity)
	battle:Dmg({Emitter=entity, Dmg=battle:RollDice(entity, {3})})
end`,
	}, nil
}

func TestFull(t *testing.T) {
	//change this
	attackSkill := "attack"
	controller, battle := getTestBattle()
	controller.SelectSkill(battle, battle.entities[0].Id, []string{attackSkill, attackSkill})
	controller.SelectSkill(battle, battle.entities[1].Id, []string{attackSkill})

	if battle.entities[0].Stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].Stats.HP)
	}
}

func TestFullWithStatus(t *testing.T) {
	//change this
	attackSkill := "attack"
	defense := "defense"
	controller, battle := getTestBattle()
	err := controller.SelectSkill(battle, battle.entities[0].Id, []string{attackSkill, attackSkill, attackSkill, attackSkill})
	if err != nil {
		t.Error(err)
	}
	err = controller.SelectSkill(battle, battle.entities[1].Id, []string{defense, attackSkill})
	if err != nil {
		t.Error(err)
	}

	if battle.entities[0].Stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 10 {
		t.Error("Expected entity 2 to have 7 HP, has", battle.entities[1].Stats.HP)
	}
}
