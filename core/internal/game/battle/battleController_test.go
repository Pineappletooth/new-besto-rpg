package battle

import "testing"

func TestFull(t *testing.T) {
	//change this
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

	selectSkill(battle, entity1.Id, []string{name})
	selectSkill(battle, entity2.Id, []string{name, name})

	if battle.entities[0].Stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].Stats.HP)
	}
	if battle.entities[1].Stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].Stats.HP)
	}
}
