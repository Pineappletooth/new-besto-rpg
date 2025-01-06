package battle

import (
	"math/rand/v2"
)

type Util struct {}

type DiceType string

const (
	D4 DiceType = "D4"
	D6 DiceType = "D6"
)


var dices = map[DiceType] []int {
	D4: {1, 2, 3, 4},
	D6: {1, 2, 3, 4, 5, 6},
}

func (s *Util) RollDice(dice []int) int{
	return dice[rand.IntN(len(dice))]
}
func (s *Util) RollDiceT(dice DiceType) int{
	return s.RollDice(dices[dice])
}