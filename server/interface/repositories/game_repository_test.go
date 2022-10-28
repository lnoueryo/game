package repositories_test

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"

	pb "game/infrastructure/grpc/proto"

	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB

// 不整合が起きていないか確認。
func TestCreateTableTransaction(t *testing.T) {
	const USER_ID = 1
	const TABLE_LIMIT = 1
	fmt.Println("Hello")
	setup(t)
	var first int
	if err := DB.QueryRow("SELECT COUNT( * ) FROM `table` WHERE adminId = ?;", USER_ID).Scan(&first); err != nil {
		t.Fatal(err)
	}
	if first != 0 {
		t.Errorf("count( * ) of table must be 0")
		return
	}
	conn := gRPCClient(t)
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = pb.NewGameClient(conn)

	requests := []*pb.TableCreateRequest{}
	for i := 0; i < 4; i++ {
		requests = append(requests, createRequest(USER_ID))
	}
	wg := &sync.WaitGroup{}  // WaitGroupの値を作る
	for i := 0; i < 20; i++ { // （例として）10回繰り返す
		wg.Add(1)  // wgをインクリメント
		go func() {
			for _, v := range requests {
				createTable(v)
			}
			wg.Done()  // 完了したのでwgをデクリメント
		}()
	}
	wg.Wait()  // メインのgoroutineはサブgoroutine 10個が完了するのを待つ
	var got int
	if err := DB.QueryRow("SELECT COUNT( * ) FROM `table` WHERE adminId = ?;", USER_ID).Scan(&got); err != nil {
		t.Fatal(err)
	}

	if got != TABLE_LIMIT {
		t.Errorf("count( * ) of table:%d, want %d", got, TABLE_LIMIT)
		return
	}

}

// func TestCreateTable(t *testing.T) {
	
// }

