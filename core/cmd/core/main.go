package main

import (
	"fmt"
	"log"
	"net"
	"pineappletooth/bestoRpg/internal/game/battle"
	"pineappletooth/bestoRpg/internal/persistence"

	"pineappletooth/bestoRpg/internal/handlers"
	pb "pineappletooth/bestoRpg/pkg/api/proto"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(recovery.UnaryServerInterceptor()),
	)
	pb.RegisterCommandsServer(s, handlers.NewCommandServer(persistence.Skill{}, persistence.Status{}, battle.MockPersistence{}))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
