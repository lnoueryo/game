package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	pb "game/table"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGameServer
}

var DB *sql.DB

// SayHello implements helloworld.GreeterServer
func (s *server) GetGame(ctx context.Context, in *pb.TableRequest) (*pb.TableReply, error) {

	rows, err := DB.Query("SELECT g.id, g.name, g.orignal_extra_fields, CASE WHEN g.key is NULL THEN JSON_ARRAY() ELSE JSON_ARRAYAGG((JSON_OBJECT('key', g.key, 'title', g.title, 'gameId', g.gameId, 'adminId', g.adminId, 'limit', g.limit, 'start', g.start, 'extraFields', g.`extraFields`, 'players', g.players))) END as tables FROM (SELECT g.id, g.name, g.extraFields AS orignal_extra_fields, t.key, t.title, t.gameId, t.adminId, t.limit, t.start, t.`extraFields`, JSON_ARRAYAGG(JSON_OBJECT('id', p.id, 'username', p.username)) as players FROM game g LEFT JOIN `table` AS t ON t.gameId = g.id LEFT JOIN player AS p ON p.tableId = t.`key` WHERE g.id = ? GROUP BY p.tableId) AS g", in.GetId())
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var tableReply pb.TableReply
	for rows.Next() {
		var extraFields []uint8
		var tables []uint8
		err := rows.Scan(&tableReply.Id, &tableReply.Name, &extraFields, &tables);if err != nil {
			fmt.Println(err)
		}
		if tableReply.Id == 0 {
			fmt.Println("Not Found")
			break;
		}
		if err := json.Unmarshal(extraFields, &tableReply.ExtraFields); err != nil {
			fmt.Println(err)
		}
		if err := json.Unmarshal(tables, &tableReply.Tables); err != nil {
			fmt.Println(err)
		}
	}

	return &tableReply, nil
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/kartenspielen")
	if err != nil {
		panic(err)
	}
	DB = db
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	pb.RegisterGameServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func getGame() {
	rows, err := DB.Query("SELECT id, name FROM game")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id, name)
	}
}