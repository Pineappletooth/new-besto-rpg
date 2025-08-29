package battle

import (
	"pineappletooth/bestoRpg/internal/game/battle/events"
	"pineappletooth/bestoRpg/internal/game/event"
	"pineappletooth/bestoRpg/internal/game/utils"
	"testing"
)

func TestRollDice(t *testing.T) {
	battleEntity := NewBattleEntity()
	battleEntity.RollDice(utils.D6)
	a := event.EventAction[events.OnRollDiceContext]{
		OnBefore: func(before events.OnRollDiceContext, after events.OnRollDiceContext) events.OnRollDiceContext {
			after.Dice = []int{1}
			return after
		},
		OnAfter: func(before events.OnRollDiceContext, after events.OnRollDiceContext) events.OnRollDiceContext {
			if len(after.Dice) != 1 || after.Dice[0] != 1 {
				t.Errorf("after.Dice Expected 1, got %d", after.Dice)
			}
			if before.Result != 1 {
				t.Errorf("before.Result Expected 1, got %d", before.Result)
			}
			if after.Result != 1 {
				t.Errorf("after.Result Expected 1, got %d", after.Result)
			}
			after.Result = 7
			return after
		},
	}
	battleEntity.rollDiceEvent.Subscribe(a)
	res := battleEntity.RollDice(utils.D6)
	if res != 7 {
		t.Errorf("Expected 7, got %d", res)
	}
}
