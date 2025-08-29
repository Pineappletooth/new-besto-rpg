package battle

type BattleContext struct {
	self    *BattleEntity
	allies  []*BattleEntity
	enemies []*BattleEntity
}
