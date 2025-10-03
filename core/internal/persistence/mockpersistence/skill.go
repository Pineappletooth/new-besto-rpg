package mockpersistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
	"pineappletooth/bestoRpg/resources"
)

var cachedSkills []model.Skill

type Skill struct{}

func (Skill) GetSkill(skill string) (model.Skill, error) {
	var skillModel model.Skill

	if cachedSkills == nil {
		var skillModelList []model.Skill
		err := json.Unmarshal(resources.Skills, &skillModelList)
		if err != nil {
			return skillModel, err
		}
		cachedSkills = skillModelList
	}
	for _, skillModelItem := range cachedSkills {
		if skillModelItem.Name == skill {
			return skillModelItem, nil
		}
	}
	return skillModel, errors.New("el skill no existe")
}
