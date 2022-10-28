package repositories

import (
	"database/sql"
	"encoding/json"

	cf "game/config"
	db "game/infrastructure/database"
	pb "game/infrastructure/grpc/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DB *sql.DB = db.DB

func GetGame(in *pb.GameRequest) (*pb.GameReply, error) {

	gameReply, err := GetGameByID(int(in.GetId()))

	return gameReply, err
}

func CreateTable(in *pb.TableCreateRequest) (*pb.GameReply, error) {
	var gameReply *pb.GameReply
	tb := in.GetTable()

	tx, _ := DB.Begin()

	err := func() error {

		// playerの共有ロックとtableIdの確認
		tableId := new(sql.NullString)
		err := DB.QueryRow("SELECT tableId FROM player WHERE id = ? FOR UPDATE", tb.AdminId).Scan(tableId);if err != nil {
			cf.Errorlog.Println(err)
			return err
		}

		if tableId.Valid {
			cf.Errorlog.Println(err)
			message := "tableId exists"
			err = status.Error(codes.Unknown, message)
			return err
		}

		// INSERT
		ins, err := tx.Prepare("INSERT `table` (`key`, title, gameId, adminId, `limit`, `start`, extraFields) VALUES(?, ?, ?, ?, ?, ?, ?)");if err != nil {
			cf.Errorlog.Println(err)
			err = status.Error(codes.Unknown, err.Error())
			return err
		}
		defer ins.Close()

		jsonObj, err := json.Marshal(tb.ExtraFields);if err != nil {
			cf.Errorlog.Println(err)
			err = status.Error(codes.Unknown, err.Error())
			return err
		}
		_, err = ins.Exec(tb.Key, tb.Title, tb.GameId, tb.AdminId, tb.Limit, tb.Start, jsonObj);if err != nil {
			cf.Errorlog.Println(err)
			err = status.Error(codes.Unknown, err.Error())
			return err
		}

		// playerのtableIdを更新
		upd, err := tx.Prepare("UPDATE player SET tableId = ? WHERE id = ?");if err != nil {
			cf.Errorlog.Println(err)
			return err
		}

		_, err = upd.Exec(tb.Key, tb.AdminId);if err != nil {
			cf.Errorlog.Println(err)
			err = status.Error(codes.Unknown, err.Error())
			return err
		}

		defer upd.Close()

		return nil
	}()

	if err != nil {
		tx.Rollback()
		message := "failed create"
		if err.Error() == "tableId exists" {
			message = "tableId exists"
		}
		err = status.Error(codes.Unknown, message)
		return gameReply, err
	} else {
		tx.Commit()
	}

	gameReply, err = GetGameByID(int(tb.GameId));if err != nil {
		cf.Errorlog.Println(err)
	}

	return gameReply, err
}

func GetGameByID(id int) (*pb.GameReply, error) {
	var gameReply pb.GameReply
	var extraFields []uint8
	var tables []uint8
	nullId := new(sql.NullInt32)
	nullName := new(sql.NullString)
	err := DB.QueryRow("SELECT g.id, g.name, g.orignal_extra_fields, CASE WHEN g.key is NULL THEN JSON_ARRAY() ELSE JSON_ARRAYAGG((JSON_OBJECT('key', g.key, 'title', g.title, 'gameId', g.gameId, 'adminId', g.adminId, 'limit', g.limit, 'start', g.start, 'extraFields', g.`extraFields`, 'players', g.players))) END as tables FROM (SELECT g.id, g.name, g.extraFields AS orignal_extra_fields, t.key, t.title, t.gameId, t.adminId, t.limit, t.start, t.`extraFields`, JSON_ARRAYAGG(JSON_OBJECT('id', p.id, 'username', p.username)) as players FROM game g LEFT JOIN `table` AS t ON t.gameId = g.id LEFT JOIN player AS p ON p.tableId = t.`key` WHERE g.id = ? GROUP BY t.`key`, p.tableId) AS g GROUP BY g.`key`, g.id", id).Scan(nullId, nullName, &extraFields, &tables);if err != nil {
		cf.Errorlog.Println(err)
	}
    if nullId.Valid {
		gameReply.Id = nullId.Int32
        gameReply.Name = nullName.String
		if err := json.Unmarshal(extraFields, &gameReply.ExtraFields); err != nil {
			cf.Errorlog.Println(err)
		}
		if err := json.Unmarshal(tables, &gameReply.Tables); err != nil {
			cf.Errorlog.Println(err)
		}
	} else {
		err = status.Error(codes.Unknown, "not found")
		cf.Errorlog.Println(err)
	}
	cf.Infolog.Println(gameReply)
	return &gameReply, err
	// rows, err := DB.Query("SELECT g.id, g.name, g.orignal_extra_fields, CASE WHEN g.key is NULL THEN JSON_ARRAY() ELSE JSON_ARRAYAGG((JSON_OBJECT('key', g.key, 'title', g.title, 'gameId', g.gameId, 'adminId', g.adminId, 'limit', g.limit, 'start', g.start, 'extraFields', g.`extraFields`, 'players', g.players))) END as tables FROM (SELECT g.id, g.name, g.extraFields AS orignal_extra_fields, t.key, t.title, t.gameId, t.adminId, t.limit, t.start, t.`extraFields`, JSON_ARRAYAGG(JSON_OBJECT('id', p.id, 'username', p.username)) as players FROM game g LEFT JOIN `table` AS t ON t.gameId = g.id LEFT JOIN player AS p ON p.tableId = t.`key` WHERE g.id = ? GROUP BY p.tableId) AS g", id);if err != nil {
	// 	cf.Errorlog.Println(err)
	// 	return &gameReply, err
	// }

	// defer rows.Close()
	// if err = rows.Err(); err != nil {
	// 	cf.Errorlog.Println(err)
	// }
	// for rows.Next() {
	// 	var extraFields []uint8
	// 	var tables []uint8
	// 	err := rows.Scan(&gameReply.Id, &gameReply.Name, &extraFields, &tables);if err != nil {
	// 		cf.Errorlog.Println(err)
	// 	}
	// 	if gameReply.Id == 0 {
	// 		fmt.Println("Not Found")
	// 		break;
	// 	}
	// 	if err := json.Unmarshal(extraFields, &gameReply.ExtraFields); err != nil {
	// 		cf.Errorlog.Println(err)
	// 	}
	// 	if err := json.Unmarshal(tables, &gameReply.Tables); err != nil {
	// 		cf.Errorlog.Println(err)
	// 	}
	// }

	// return &gameReply, err
}