package battle

import (
	"github.com/google/uuid"
)

type Battle struct {
	id    uuid.UUID
	team1 []BattleEntity
	team2 []BattleEntity
}

func initBattle(team1 []BattleEntity, team2 []BattleEntity) {
	battle := &Battle{
		id:    uuid.New(),
		team1: team1,
		team2: team2,
	}
}

func startRound(battle *Battle) {
	index := 0
	for {
		for _, entity := range battle.team1 {
			if len(entity.chosenSkills) > 0 {
				entity.chosenSkills[index]
			}
		}
	}
}
