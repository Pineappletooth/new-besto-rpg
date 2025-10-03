package handlers

import (
	"pineappletooth/bestoRpg/internal/game/battle"
	pb "pineappletooth/bestoRpg/pkg/api/proto"
)

type commandServer struct {
	pb.UnimplementedCommandsServer
	SkillPersistence  battle.SkillPersistence
	StatusPersistence battle.StatusPersistence
	BattlePersistence battle.Persistence
}

func NewCommandServer(skillPersistence battle.SkillPersistence, statusPersistence battle.StatusPersistence, battlePersistence battle.Persistence) pb.CommandsServer {
	return &commandServer{
		SkillPersistence:  skillPersistence,
		StatusPersistence: statusPersistence,
		BattlePersistence: battlePersistence,
	}
}
