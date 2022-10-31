package grpc

import (
	"log"
	"net"
	"os"

	pb "game/infrastructure/grpc/proto"
	cr "game/interface/controllers"
	"game/config"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)


func CreateServer() error {
	SERVER_PORT := os.Getenv("SERVER_PORT");if SERVER_PORT == "" {
		SERVER_PORT = "50051"
	}

	lis, err := net.Listen("tcp", ":" + SERVER_PORT);if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	middleware := config.CreateAccesslog()

	s := grpc.NewServer(
		middleware,
	)

	pb.RegisterGameServer(s, &cr.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return err
}
