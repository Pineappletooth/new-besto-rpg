package battle

import (
	"pineappletooth/bestoRpg/internal/game/utils"
	"slices"
)

var skills = make(map[string]Skill)

type Battle struct {
	Id       string
	entities []BattleEntity
}

type dmgCtx struct {
	Emitter *BattleEntity
	Target  *BattleEntity
	Dmg     int
}

func (battle *Battle) UseSkill(name string, selfEntity *BattleEntity) {
	skill, ok := battle.GetSkill(name)
	if !ok {
		return
	}
	//TODO: Event on cast skill
	skill.OnUse(battle, selfEntity)
	//TODO: Event after cast skill
}

func (*Battle) GetSkill(name string) (Skill, bool) {
	skill, ok := skills[name]
	return skill, ok
}

func (battle *Battle) Dmg(ctx dmgCtx) {
	target := ctx.Target
	if target == nil {
		target = battle.getTarget(ctx.Emitter.Team)
	}

	//on before event

	target.Stats.HP = target.Stats.HP - ctx.Dmg

	//on after event
}

func (battle *Battle) RollDice(emitter *BattleEntity, dice []int) int {

	before := emitter.Events.onBeforeRollDice.Emit(onBeforeRollDiceContext{
		dice,
	})

	result := utils.RollDice(before.Dice)

	after := emitter.Events.onAfterRollDice.Emit(onAfterRollDiceContext{
		dice,
		result,
	})
	return after.Result

}

func (battle *Battle) getEnemies(team int) []*BattleEntity {
	enemies := make([]*BattleEntity, 0, len(battle.entities)-1)
	for i := range battle.entities {
		if team != battle.entities[i].Team {
			enemies = append(enemies, &battle.entities[i])
		}
	}
	return enemies
}

func (battle *Battle) getTarget(team int) *BattleEntity {
	enemies := battle.getEnemies(team)
	slices.SortFunc(enemies, func(e, e2 *BattleEntity) int {
		return e.Stats.Aggro - e2.Stats.Aggro
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
