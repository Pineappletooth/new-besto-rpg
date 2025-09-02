package battle

import (
	"pineappletooth/bestoRpg/internal/game/event"
	"pineappletooth/bestoRpg/internal/game/utils"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type Battle struct {
	id       uuid.UUID
	entities []BattleEntity
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

func (battle *Battle) processRound() {
	for _, entity := range battle.entities {
		for _, skillName := range entity.chosenSkills {
			entity.skills[skillName].onUse(battle, &entity)
		}
	}
}

func (battle *Battle) dmg(emitter *BattleEntity, dmg int) {
	target := battle.getTarget(emitter.team)

	//on before event

	target.stats.HP = target.stats.HP - dmg

	//on after event
}

func (battle *Battle) rollDice(emitter *BattleEntity, dice []int) int {

	before := emitter.events.onBeforeRollDice.Emit(onBeforeRollDiceContext{
		dice,
	})

	result := utils.RollDice(before.Dice)

	after := emitter.events.onAfterRollDice.Emit(onAfterRollDiceContext{
		dice,
		result,
	})
	return after.Result

}

func (battle *Battle) getEnemies(team int) []*BattleEntity {
	enemies := make([]*BattleEntity, 0, len(battle.entities)-1)
	for _, battleEntity := range battle.entities {
		if team != battleEntity.team {
			enemies = append(enemies, &battleEntity)
		}
	}
	return enemies
}

func (battle *Battle) getTarget(team int) *BattleEntity {
	enemies := battle.getEnemies(team)
	slices.SortFunc(enemies, func(e, e2 *BattleEntity) int {
		return e.stats.Aggro - e2.stats.Aggro
	})
	print(enemies)
	return enemies[len(enemies)-1]
}
