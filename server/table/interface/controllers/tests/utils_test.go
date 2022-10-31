package controllers_test

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"testing"

	pb "game/infrastructure/grpc/proto"
	db "game/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DB *sql.DB

func init() {
	DB = db.DB
}

func setup(t *testing.T) {
	if _, err := DB.Exec("DELETE FROM `table`"); err != nil {
		t.Fatal(err)
	}
	// 空であることを確認
	var first int
	if err := DB.QueryRow("SELECT COUNT( * ) FROM `table`").Scan(&first); err != nil {
		t.Fatal(err)
	}
	if first != 0 {
		t.Errorf("count( * ) of table must be 0")
		return
	}
}

var (
	scanner *bufio.Scanner
	client  pb.GameClient
)

func gRPCClient(t *testing.T) *grpc.ClientConn {
	// fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	);if err != nil {
		t.Fatal("Connection failed.")
		return conn
	}

	return conn

}

func createRequest(adminId int32, GameId int32) *pb.TableCreateRequest {
	u, err := uuid.NewRandom();if err != nil {
		fmt.Println(err)
	}
	uu := u.String()
	table := pb.Table{
		Key:         uu,
		Title:       "Hello",
		GameId:      GameId,
		AdminId:     adminId,
		Limit:       3,
		ExtraFields: map[string]string{"turn": ""},
	}
	return &pb.TableCreateRequest{
		Table: &table,
	}
}
