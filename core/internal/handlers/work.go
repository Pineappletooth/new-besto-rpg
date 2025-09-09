package handlers

import (
	"context"
	"pineappletooth/bestoRpg/internal/persistence"
	pb "pineappletooth/bestoRpg/pkg/api/proto"

	"google.golang.org/protobuf/proto"
)

var com = command{
	name:            "work",
	cooldownSeconds: 10,
}

func (s *commandServer) Work(context context.Context, request *pb.WorkRequest) (*pb.WorkResponse, error) {
	err := com.validateAll(request.GetUserId())
	if err != nil {
		return nil, err
	}
	defer com.postCommand(request.GetUserId())

	char, err := persistence.GetCharacter(request.GetUserId())
	if err != nil {
		return nil, err
	}

	char.Inventory.Gold += 100
	err = persistence.AddCharacter(char)
	if err != nil {
		return nil, err
	}
	return &pb.WorkResponse{
		Amount:     proto.Int64(100),
		NewBalance: proto.Int64(char.Inventory.Gold),
	}, nil
}
