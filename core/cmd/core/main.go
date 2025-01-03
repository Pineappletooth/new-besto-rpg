package main

import (
	"fmt"
	"log"
	"net"

	"pineappletooth/bestoRpg/internal/handlers"
	pb "pineappletooth/bestoRpg/pkg/api/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommandsServer(s, &handlers.CommandServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
