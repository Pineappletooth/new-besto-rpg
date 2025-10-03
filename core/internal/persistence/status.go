package persistence

import (
	"encoding/json"
	"errors"
	"pineappletooth/bestoRpg/internal/model"
)

type Status struct{}

func (Status) AddStatus(status model.Status) error {
	res := redisClient.JSONSet(ctx, "status:"+status.Name, "$", status)
	return res.Err()
}

func (Status) GetStatus(status string) (model.Status, error) {
	if redisClient.Exists(ctx, "status:"+status).Val() == 0 {
		return model.Status{}, errors.New("el status no existe")
	}
	res := redisClient.JSONGet(ctx, "status:"+status)
	if res.Err() != nil {
		return model.Status{}, res.Err()
	}
	var character model.Status
	err := json.Unmarshal([]byte(res.Val()), &character)
	return character, err
}
