package resources

import _ "embed"

//go:embed battleEntity.json
var BattleEntity []byte

//go:embed skills.json
var Skills []byte

//go:embed status.json
var Status []byte
