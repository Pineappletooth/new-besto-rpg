package battle

import (
	"pineappletooth/bestoRpg/internal/game/event"
	"strings"

	"github.com/google/uuid"
)

type Battle struct {
	id    uuid.UUID
	team1 []BattleEntity
	team2 []BattleEntity
}

type Events struct {
	onBeforeRollDice event.Event[onBeforeRollDiceContext]
	onAfterRollDice  event.Event[onAfterRollDiceContext]
}

func newEvents() Events {
	return Events{
		onBeforeRollDice: event.New[onBeforeRollDiceContext](),
		onAfterRollDice:  event.New[onAfterRollDiceContext](),
	}
}

func parseCommandMock(first string) {
	strings.Split(first, " ")
}

func (battle *Battle) startRound() {
	for _, entity := range battle.team1 {
		for _, skillName := range entity.chosenSkills {
			entity.skills[skillName].onUse()
		}
	}
}

func processRound()
