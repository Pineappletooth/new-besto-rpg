package battle

import (
	"github.com/google/uuid"
)

type Battle struct {
	id     uuid.UUID
	team1  []BattleEntity
	team2  []BattleEntity
	events Events
}

type Events struct {
	onRollDice Event[onRollDiceContext]
}

func initBattle(team1 []BattleEntity, team2 []BattleEntity) {

}

func startRound(battle *Battle) {
	for {
		for _, entity := range battle.team1 {
			if len(entity.chosenSkills) > 0 {
			}
		}
	}
}
