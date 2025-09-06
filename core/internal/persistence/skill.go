package persistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
)

func AddSkill(skill model.Skill) error {
	res := redisClient.JSONSet(ctx, "skill:"+skill.Name, "$", skill)
	return res.Err()
}

func GetSkill(skill string) (model.Skill, error) {
	if redisClient.Exists(ctx, "skill:"+skill).Val() == 0 {
		return model.Skill{}, errors.New("el skill no existe")
	}
	res := redisClient.JSONGet(ctx, "skill:"+skill)
	if res.Err() != nil {
		return model.Skill{}, res.Err()
	}
	var character model.Skill
	err := json.Unmarshal([]byte(res.Val()), &character)
	return character, err
}
