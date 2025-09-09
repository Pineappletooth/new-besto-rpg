package handlers

import (
	"context"
	"pineappletooth/bestoRpg/internal/game/battle"
	"pineappletooth/bestoRpg/internal/persistence"
	pb "pineappletooth/bestoRpg/pkg/api/proto"
)

var pvpBattleCommand = command{
	name:            "work",
	cooldownSeconds: 10,
}

func (s *commandServer) PvpBattle(context context.Context, request *pb.PvpBattleRequest) (*pb.PvpBattleResponse, error) {
	/*err := pvpBattleCommand.validateAll(request.GetUserId())
	if err != nil {
		return nil, err
	}
	defer com.postCommand(request.GetUserId())
	*/
	char, err := persistence.GetCharacter(request.GetUserId())
	if err != nil {
		return nil, err
	}
	char2, err := persistence.GetCharacter(request.GetOpponentId())
	if err != nil {
		return nil, err
	}

	controller := battle.Controller{
		SkillPersistence:  s.SkillPersistence,
		BattlePersistence: s.BattlePersistence,
	}

	entity := battle.NewBattleEntityFromCharacter(char)
	entity2 := battle.NewBattleEntityFromCharacter(char2)
	newBattle, err := battle.NewBattle(controller, []battle.BattleEntity{entity, entity2})
	if err != nil {
		return nil, err
	}

	return &pb.PvpBattleResponse{
		BattleId: &newBattle.Id,
	}, nil
}
