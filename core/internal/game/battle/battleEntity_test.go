package battle

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRollDice(t *testing.T) {
	battleEntity := NewBattleEntity()
	battleEntity.RollDice(D6)
	a := EventAction[onRollDiceContext]{
		onBeforeEvent: func(ctx EventContext[onRollDiceContext]) {
			fmt.Print("BBBB")
			ctx.modified.result = 4
		},
		onAfterEvent: func(ctx EventContext[onRollDiceContext]) {
			fmt.Print("after", ctx.modified.result)
			ctx.modified.result = 7
		},
	}
	battleEntity.rollDiceEvent.subscribe(a)
	res := battleEntity.RollDice(D6)
	if res != 7 {
		t.Errorf("Expected 7, got %d", res)
	}
}
