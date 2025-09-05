package battle

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"pineappletooth/bestoRpg/internal/model"
)

type OnUseSkill func(battle *Battle, self *BattleEntity)
type Skill struct {
	name  string
	OnUse OnUseSkill
}

func NewSkillFromModel(dto model.Skill) *Skill {
	L := lua.NewState()
	defer L.Close()
	skill := &Skill{
		name: dto.Name,
	}
	L.SetGlobal("skill", luar.New(L, skill))
	if err := L.DoString(dto.Action); err != nil {
		panic(err)
	}
	if skill.OnUse == nil {
		panic("no onUse")
	}
	return skill
}

type Status struct {
	name     string
	priority int
	//effect   *Effect[Context]
}
