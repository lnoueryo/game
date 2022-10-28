package repositories_test

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	pb "game/infrastructure/grpc/proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DATABASE_URL = "root:password@tcp(127.0.0.1:3307)/kartenspielen"
)

func setup(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/kartenspielen")
	if err != nil {
		t.Fatal(err)
	}
	DB = db
	// クリーンアップ
	if _, err := DB.Exec("DELETE FROM `table`"); err != nil {
		t.Fatal(err)
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
	)
	if err != nil {
		t.Fatal("Connection failed.")
		return conn
	}

	return conn

}

func getGame() *pb.GameReply {

	req := &pb.GameRequest{
		Id: 1,
	}
	res, err := client.GetGame(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	return res
}

func createTable(req *pb.TableCreateRequest) *pb.GameReply {

	res, _ := client.CreateTable(context.Background(), req)

	return res
}

func createRequest(adminId int32) *pb.TableCreateRequest {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	uu := u.String()
	table := pb.Table{
		Key:         uu,
		Title:       "Hello",
		GameId:      1,
		AdminId:     adminId,
		Limit:       3,
		Start:       0,
		ExtraFields: map[string]string{"turn": ""},
	}
	return &pb.TableCreateRequest{
		Table: &table,
	}
}
