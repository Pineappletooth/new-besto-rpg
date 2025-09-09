package battle

import "fmt"

var battlePersitence = make(map[string]*Battle)

type MockPersistence struct {
}

func (m MockPersistence) SaveBattle(battle *Battle) error {
	battlePersitence[battle.Id] = battle
	return nil
}

func (m MockPersistence) GetBattle(battleId string) (*Battle, error) {
	battle, ok := battlePersitence[battleId]
	if !ok {
		return nil, fmt.Errorf("battle not found")
	}
	return battle, nil
}
