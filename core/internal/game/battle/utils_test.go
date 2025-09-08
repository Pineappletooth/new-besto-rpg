package battle

func getTestBattle() (battleController, *Battle) {
	entity1 := NewBattleEntityTest()
	entity2 := NewBattleEntityTest()

	entity1.Team = 1
	entity2.Team = 2

	controller := battleController{
		skillPersistence: mockSkillPersistence{},
	}

	battle := NewBattle(controller, []BattleEntity{entity1, entity2})
	return controller, battle
}
