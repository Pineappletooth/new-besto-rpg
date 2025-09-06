package battle

import (
	"pineappletooth/bestoRpg/internal/model"
	"testing"
)

func TestNewSkill(t *testing.T) {
	dto := model.Skill{
		Name: "test",
		Action: `
skill.OnUse = function (battle, entity)
	battle:Dmg({Emitter=entity, Dmg=battle:RollDice(entity, {3})})
end`,
	}

	s := NewSkillFromModel(dto)

	entity1 := NewBattleEntityTest()
	entity2 := NewBattleEntityTest()
	name := "attack"
	skill := Skill{
		name,
		func(battle *Battle, self *BattleEntity) {
			battle.Dmg(dmgCtx{Emitter: self, Dmg: battle.RollDice(self, []int{3})})
		},
	}

	skills[name] = skill

	entity1.Team = 1
	entity2.Team = 2

	battle := New([]BattleEntity{entity1, entity2})
	e, _ := battle.getEntityById(entity1.Id)
	println(battle.Id)
	s.OnUse(battle, e)
	if battle.entities[1].Stats.HP != 7 {
		t.Error("Expected entity 2 to have 7 HP, has", battle.entities[1].Stats.HP)
	}
}
