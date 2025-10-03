package persistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
)

type Skill struct{}

func (Skill) AddSkill(skill model.Skill) error {
	res := redisClient.JSONSet(ctx, "skill:"+skill.Name, "$", skill)
	return res.Err()
}

func (Skill) GetSkill(skill string) (model.Skill, error) {
	if redisClient.Exists(ctx, "skill:"+skill).Val() == 0 {
		return model.Skill{}, errors.New("el skill no existe")
	}
	res := redisClient.JSONGet(ctx, "skill:"+skill)
	if res.Err() != nil {
		return model.Skill{}, res.Err()
	}
	var skillModel model.Skill
	err := json.Unmarshal([]byte(res.Val()), &skillModel)
	return skillModel, err
}
