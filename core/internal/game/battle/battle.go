package battle

import (
	"fmt"
	"pineappletooth/bestoRpg/internal/game/utils"
	"slices"
)

type Battle struct {
	Id       string
	entities []*BattleEntity
	skills   map[string]*Skill
	status   map[string]*Status
}

type dmgCtx struct {
	Emitter *BattleEntity
	Target  *BattleEntity
	Dmg     int
}

type statusCtx struct {
	Emitter *BattleEntity
	Target  *BattleEntity
	Status  string
	Turns   int
}

func (battle *Battle) ApplyStatus(ctx statusCtx) {
	if ctx.Emitter == nil {
		fmt.Println("Must send emitter")
		return
	}
	if ctx.Turns == 0 {
		fmt.Println("Must send turns")
	}
	if ctx.Status == "" {
		fmt.Println("Must send status")
	}

	target := ctx.Target
	if target == nil {
		var err error
		target, err = battle.getTarget(ctx.Emitter.Team)
		if err != nil {
			fmt.Printf("no target %v", err)
			return
		}
	}
	//TODO: Lazy load status
	status := battle.GetStatus(ctx.Status)

	target.Status = append(target.Status, BattleStatus{
		Status:   *status,
		Duration: ctx.Turns,
	})
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

func (battle *Battle) GetSkill(name string) (*Skill, bool) {
	skill, ok := battle.skills[name]
	return skill, ok
}

func (battle *Battle) GetStatus(name string) *Status {
	status, ok := battle.status[name]
	if !ok {
		battle.status[name] = &Status{
			Name:    name,
			OnApply: nil,
		}
		return battle.status[name]
	}
	return status
}

func (battle *Battle) Dmg(ctx dmgCtx) {
	if ctx.Emitter == nil {
		fmt.Println("Must send emitter")
		return
	}
	if ctx.Dmg == 0 {
		fmt.Println("Must send dmg")
		return
	}

	target := ctx.Target
	if target == nil {
		var err error
		target, err = battle.getTarget(ctx.Emitter.Team)
		if err != nil {
			fmt.Printf("no target %v", err)
			return
		}
	}

	res := target.Events.OnBeforeDmg.Emit(OnBeforeDmgContext{
		Emitter: ctx.Emitter,
		Target:  target,
		Dmg:     ctx.Dmg,
	})
	ctx.Emitter = res.Emitter
	ctx.Target = res.Target
	ctx.Dmg = res.Dmg

	target.Stats.HP = target.Stats.HP - ctx.Dmg

	//on after event
}

func (battle *Battle) RollDice(emitter *BattleEntity, dice []int) int {

	before := emitter.Events.OnBeforeRollDice.Emit(OnBeforeRollDiceContext{
		dice,
	})

	result := utils.RollDice(before.Dice)

	after := emitter.Events.OnAfterRollDice.Emit(OnAfterRollDiceContext{
		dice,
		result,
	})
	return after.Result

}

func (battle *Battle) getEnemies(team int) []*BattleEntity {
	enemies := make([]*BattleEntity, 0, len(battle.entities)-1)
	for i := range battle.entities {
		entity := battle.entities[i]
		if team != entity.Team || -1 == entity.Team {
			enemies = append(enemies, entity)
		}
	}
	return enemies
}

func (battle *Battle) getTarget(team int) (*BattleEntity, error) {
	enemies := battle.getEnemies(team)
	if len(enemies) == 0 {
		return nil, fmt.Errorf("no enemies")
	}
	slices.SortFunc(enemies, func(e, e2 *BattleEntity) int {
		return e.Stats.Aggro - e2.Stats.Aggro
	})
	print(enemies)
	return enemies[len(enemies)-1], nil
}

func (battle *Battle) getEntityById(id string) (*BattleEntity, bool) {
	for i := range battle.entities {
		if battle.entities[i].Id == id {
			return battle.entities[i], true
		}
	}
	return nil, false
}
