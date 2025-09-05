package main

import (
	"pineappletooth/bestoRpg/internal/model"
	"pineappletooth/bestoRpg/internal/persistence"
)

func main() {
	err := persistence.AddCharacter(model.Character{
		Id: "0",
		Inventory: model.Inventory{
			Items: map[string]int{
				"potion": 5,
				"scroll": 3,
			},
			Gold: 100,
		},
	})
	if err != nil {
		println(err.Error())
	}
}
