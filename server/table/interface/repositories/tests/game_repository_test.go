package repositories_test

import (
	"sync"
	"testing"

	pb "game/infrastructure/grpc/proto"
	repo "game/interface/repositories"
	_ "github.com/go-sql-driver/mysql"
)

// 不整合が起きていないか確認。
func TestCreateTableTransaction(t *testing.T) {
	const USER_ID = 1
	const GAME_ID = 1
	const TABLE_LIMIT = 1

	setup(t)

	requests := []*pb.TableCreateRequest{}
	for i := 0; i < 4; i++ {
		requests = append(requests, createRequest(USER_ID, GAME_ID))
	}

	wg := &sync.WaitGroup{}  // WaitGroupの値を作る
	for i := 0; i < 20; i++ {
		wg.Add(1)  // wgをインクリメント
		go func() {
			for _, v := range requests {
				repo.CreateTable(v)
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

func TestCreateTable(t *testing.T) {
	const PLAYER_NUM = 11
	const GAME_ID = 1
	const TABLE_LIMIT = 1

	setup(t)

	rows, err := DB.Query("SELECT id FROM player");if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	ids := []int{}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			t.Fatalf("getRows rows.Scan error err:%v", err)
		}
		ids = append(ids, id)
	}

	err = rows.Err()
	if err != nil {
		t.Fatalf("getRows rows.Err error err:%v", err)
	}

	if len(ids) != PLAYER_NUM {
		t.Errorf("count( * ) of player must be %d", PLAYER_NUM)
	}

	requests := []*pb.TableCreateRequest{}
	for _, v := range ids {
		id := int32(v)
		requests = append(requests, createRequest(id, GAME_ID))
	}

	wg := &sync.WaitGroup{}  // WaitGroupの値を作る
	for i := 0; i < 5; i++ {
		wg.Add(1)  // wgをインクリメント
		go func() {
			for _, v := range requests {
				repo.CreateTable(v)
			}
			wg.Done()  // 完了したのでwgをデクリメント
		}()
	}
	wg.Wait()  // メインのgoroutineはサブgoroutine 10個が完了するのを待つ

	// 作成されたテーブルの合計数
	var got int
	if err := DB.QueryRow("SELECT COUNT( * ) FROM `table`").Scan(&got); err != nil {
		t.Fatal(err)
	}

	if got != PLAYER_NUM {
		t.Errorf("count( * ) of table:%d, want %d", got, PLAYER_NUM)
		return
	}

	for _, v := range requests {

		var count int
		var tableId string
		var adminId = v.GetTable().AdminId
		if err := DB.QueryRow("SELECT COUNT( * ), `key` FROM `table` WHERE adminId = ?;", adminId).Scan(&count, &tableId); err != nil {
			t.Fatal(err)
		}

		if count != TABLE_LIMIT {
			t.Errorf("count( * ) of table:%d, want %d", got, TABLE_LIMIT)
			return
		}

		var isPlayer int
		if err := DB.QueryRow("SELECT EXISTS(SELECT * FROM player WHERE tableId = ? AND id = ?) AS player;", tableId, v.GetTable().AdminId).Scan(&isPlayer); err != nil {
			t.Fatal(err)
		}

		if isPlayer == 0 {
			t.Errorf("playerid %d doesn't have the table %s", adminId, tableId)
		}

	}

}
