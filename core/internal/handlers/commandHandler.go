package handlers

import (
	"pineappletooth/bestoRpg/internal/game/battle"
	pb "pineappletooth/bestoRpg/pkg/api/proto"
)

type commandServer struct {
	pb.UnimplementedCommandsServer
	SkillPersistence  battle.SkillPersistence
	BattlePersistence battle.Persistence
}

func NewCommandServer(skillPersistence battle.SkillPersistence, battlePersistence battle.Persistence) pb.CommandsServer {
	return &commandServer{
		SkillPersistence:  skillPersistence,
		BattlePersistence: battlePersistence,
	}
}
