package battle

import (
	"pineappletooth/bestoRpg/internal/game/battle/events"
	"pineappletooth/bestoRpg/internal/game/event"

	"github.com/google/uuid"
)

type Battle struct {
	id     uuid.UUID
	team1  []BattleEntity
	team2  []BattleEntity
	events Events
}

type Events struct {
	onRollDice event.Event[events.OnRollDiceContext]
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
