package battle

import (
	"pineappletooth/bestoRpg/internal/model"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type OnUseSkill func(battle *Battle, self *BattleEntity)
type OnApplySkill func(battle *Battle, self *BattleEntity)
type Skill struct {
	name  string
	OnUse OnUseSkill
}

func NewSkillFromModel(dto model.Skill) *Skill {
	L := lua.NewState()
	defer L.Close()
	skill := &Skill{
		name:  dto.Name,
		OnUse: nil,
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
	Name    string
	OnApply OnApplySkill
}

func NewStatusFromModel(dto model.Status) *Status {
	L := lua.NewState()
	defer L.Close()
	status := &Status{
		Name:    dto.Name,
		OnApply: nil,
	}
	L.SetGlobal("status", luar.New(L, status))
	if err := L.DoString(dto.Action); err != nil {
		panic(err)
	}
	if status.OnApply == nil {
		panic("no OnApply")
	}
	return status
}
