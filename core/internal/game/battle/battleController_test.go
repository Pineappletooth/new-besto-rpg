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
			battle.dmg(dmgCtx{emitter: self, dmg: battle.rollDice(self, []int{3})})
		},
	}

	entity1.team = 1
	entity2.team = 2
	entity1.skills[name] = skill
	entity2.skills[name] = skill

	battle := New([]BattleEntity{entity1, entity2})

	selectSkill(battle, entity1.id, []string{name})
	selectSkill(battle, entity2.id, []string{name, name})

	if battle.entities[0].stats.HP != 7 {
		t.Error("Expected entity 1 to have 7 HP, has", battle.entities[0].stats.HP)
	}
	if battle.entities[1].stats.HP != 4 {
		t.Error("Expected entity 2 to have 4 HP, has", battle.entities[1].stats.HP)
	}
}
