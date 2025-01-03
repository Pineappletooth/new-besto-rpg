package persistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
	"strconv"
)

func AddCharacter(character model.Character) error {
	res := redisClient.JSONSet(ctx, "user:"+strconv.FormatUint(uint64(character.Id),10), "$",character)
	return res.Err()
}

func GetCharacter(id uint32) (model.Character, error) {
	userId := strconv.FormatUint(uint64(id),10)
	if redisClient.Exists(ctx, "user:"+userId).Val() == 0 {
		return model.Character{}, errors.New("el personaje no existe")
	}
	res := redisClient.JSONGet(ctx, "user:"+userId)
	if res.Err() != nil {
		return model.Character{}, res.Err()
	}
	var character model.Character
	err := json.Unmarshal([]byte(res.Val()), &character)
	return character, err
}
