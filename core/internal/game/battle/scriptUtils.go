package battle

import (
	"math/rand/v2"
)


type DiceType string

const (
	D4 DiceType = "d4"
	D6 DiceType = "d6"
)


var dices = map[DiceType] []int {
	D4: {1, 2, 3, 4},
	D6: {1, 2, 3, 4, 5, 6},
}

func GetDice(dice DiceType) []int{
	values := dices[dice]
	if len(values) == 0 {
		return dices[D6]
	}
	return values
}

func RollDice(dice []int) int{
	return dice[rand.IntN(len(dice))]
}