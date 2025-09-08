package main

import (
	"pineappletooth/bestoRpg/internal/model"
	"pineappletooth/bestoRpg/internal/persistence"
)

func main() {
	addChar()
	addSkill()
}

func addChar() {
	items := make(map[model.EquipmentType]model.Item)
	items[model.Head] = model.Item{
		Name:   "basic",
		Skills: []string{"attack"},
		Stats: model.Stats{
			HP:    10,
			Aggro: 0,
		},
	}

	err := persistence.AddCharacter(model.Character{
		Id: "0",
		Inventory: model.Inventory{
			Items: map[string]int{
				"potion": 5,
				"scroll": 3,
			},
			Gold: 100,
		},
		Equipment: model.Equipment{
			Items: items,
		},
	})
	if err != nil {
		println(err.Error())
	}
}

func addSkill() {
	skill := model.Skill{
		Name: "attack",
		Action: `skill.OnUse = function (battle, entity)
	battle:Dmg({Emitter=entity, Dmg=battle:RollDice(entity, {3})})
end`,
	}
	err := persistence.AddSkill(skill)
	if err != nil {
		println(err.Error())
	}
}
