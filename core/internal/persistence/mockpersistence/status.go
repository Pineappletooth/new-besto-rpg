package mockpersistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
	"pineappletooth/bestoRpg/resources"
)

var cachedStatus []model.Status

type Status struct{}

func (Status) GetStatus(status string) (model.Status, error) {
	var statusModel model.Status

	if cachedStatus == nil {
		var statusList []model.Status
		err := json.Unmarshal(resources.Skills, &statusList)
		if err != nil {
			return statusModel, err
		}
		cachedStatus = statusList
	}
	for _, statusModelItem := range cachedSkills {
		if statusModelItem.Name == status {
			return statusModel, nil
		}
	}
	return statusModel, errors.New("el status no existe")
}
