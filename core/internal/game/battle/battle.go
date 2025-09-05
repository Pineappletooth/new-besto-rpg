package battle

import (
	"pineappletooth/bestoRpg/internal/game/utils"
	"slices"
)

type Battle struct {
	Id       string
	entities []BattleEntity
}

type dmgCtx struct {
	Emitter *BattleEntity
	Target  *BattleEntity
	Dmg     int
}

func (battle *Battle) Dm2() {
	println("KK")
}

func (battle *Battle) Dmg(ctx dmgCtx) {
	target := ctx.Target
	if target == nil {
		target = battle.getTarget(ctx.Emitter.team)
	}

	//on before event

	target.stats.HP = target.stats.HP - ctx.Dmg

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
	for i := range battle.entities {
		if team != battle.entities[i].team {
			enemies = append(enemies, &battle.entities[i])
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

func (battle *Battle) getEntityById(id string) (*BattleEntity, bool) {
	for i := range battle.entities {
		if battle.entities[i].Id == id {
			return &battle.entities[i], true
		}
	}
	return nil, false
}
