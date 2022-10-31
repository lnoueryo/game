package controllers_test

import (
	"context"
	"testing"

	pb "game/infrastructure/grpc/proto"

	_ "github.com/go-sql-driver/mysql"
)


func TestCreateTable(t *testing.T) {
	const USER_ID = 1
	const GAME_ID = 1

	setup(t)

	conn := gRPCClient(t)
	defer conn.Close()

	client = pb.NewGameClient(conn)

	request := createRequest(USER_ID, GAME_ID)

	res, err := client.CreateTable(context.Background(), request);if err != nil {
		t.Error(err)
	}

	if res.Id != GAME_ID {
		t.Errorf("game id must be %d but got %d", GAME_ID, res.Id)
	}

	if res.Tables[0].AdminId != USER_ID {
		t.Errorf("table admin id must be %d but got %d", USER_ID, res.Tables[0].AdminId)
	}
}

func TestGetGame(t *testing.T) {
	const GAME_ID = 1

	conn := gRPCClient(t)
	defer conn.Close()

	client = pb.NewGameClient(conn)

	req := &pb.GameRequest{
		Id: GAME_ID,
	}
	res, err := client.GetGame(context.Background(), req);if err != nil {
		t.Error(err)
	}

	if len(res.Tables) == 0 {
		t.Error(err)
	}

	if len(res.Tables[0].Players) == 0 {
		t.Error(err)
	}

}