package controllers

import (
	"context"

	pb "game/infrastructure/grpc/proto"
	repo "game/interface/repositories"

	_ "github.com/go-sql-driver/mysql"
	// "google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedGameServer
}

func (s *Server) GetGame(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {

	gameReply, err := repo.GetGame(in)

	return gameReply, err
}

// func getGame(id int) {

// }
// SayHello implements helloworld.GreeterServer
func (s *Server) CreateTable(ctx context.Context, in *pb.TableCreateRequest) (*pb.GameReply, error) {

	gameReply, err := repo.CreateTable(in)

	return gameReply, err
}