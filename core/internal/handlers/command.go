package handlers

import (
	"pineappletooth/bestoRpg/internal/persistence"
	"time"

	pb "pineappletooth/bestoRpg/pkg/api/proto"

	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type command struct {
	name     string
	cooldownSeconds int64
}

func (com command) validateAll(userId uint32) (error) {
	return com.validateCooldown(userId)
}

func (com command) postCommand(userId uint32) {
	persistence.AddCommandLastUsed(userId, com.name)
}

func (com command) validateCooldown(userId uint32) error {
	last, err := persistence.GetCommandLastUsed(userId, com.name)
	if err != nil {
		println(err.Error())
		return err
	}
	timeNow := time.Now().Unix()

	if timeNow-last < com.cooldownSeconds {
		st := status.New(codes.FailedPrecondition, "El comando esta en cooldown, quedan "+strconv.FormatInt(com.cooldownSeconds - (timeNow - last), 10)+" segundos")
		st,err = st.WithDetails(pb.Error_builder{Details:  map[string]string{
			"cooldown": strconv.FormatInt(com.cooldownSeconds - (timeNow - last), 10),
		}}.Build())
		if err != nil {
			return status.Errorf(codes.Internal, "Error al construir el mensaje de error: %v", err)
		}
		return st.Err()
	}
	return nil
}