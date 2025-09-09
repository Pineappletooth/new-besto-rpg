package persistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
)

func AddCharacter(character model.Character) error {
	res := redisClient.JSONSet(ctx, "user:"+character.Id, "$", character)
	return res.Err()
}

func GetCharacter(userId string) (model.Character, error) {
	if redisClient.Exists(ctx, "user:"+userId).Val() == 0 {
		return model.Character{}, errors.New("character not found")
	}
	res := redisClient.JSONGet(ctx, "user:"+userId)
	if res.Err() != nil {
		return model.Character{}, res.Err()
	}
	var character model.Character
	err := json.Unmarshal([]byte(res.Val()), &character)
	return character, err
}
