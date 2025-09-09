package handlers

import (
	"context"
	"encoding/json"
	battle2 "pineappletooth/bestoRpg/internal/game/battle"
	pb "pineappletooth/bestoRpg/pkg/api/proto"
)

func (s *commandServer) SelectSkills(context context.Context, request *pb.SelectSkillsRequest) (*pb.SelectSkillsResponse, error) {
	controller := battle2.Controller{BattlePersistence: s.BattlePersistence}

	battle, err := controller.BattlePersistence.GetBattle(request.GetBattleId())
	if err != nil {
		return nil, err
	}
	err = controller.SelectSkill(battle, request.GetUserId(), request.GetSkills())
	if err != nil {
		return nil, err
	}
	str, err := json.Marshal(battle)
	if err != nil {
		return nil, err
	}
	response := string(str)
	return &pb.SelectSkillsResponse{
		Message: &response,
	}, nil
}
