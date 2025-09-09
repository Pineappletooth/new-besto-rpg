package battle

func getTestBattle() (Controller, *Battle) {
	entity1 := NewBattleEntityTest()
	entity2 := NewBattleEntityTest()

	entity1.Team = 1
	entity2.Team = 2

	controller := Controller{
		SkillPersistence: mockSkillPersistence{},
	}

	battle, _ := NewBattle(controller, []BattleEntity{entity1, entity2})
	return controller, battle
}
