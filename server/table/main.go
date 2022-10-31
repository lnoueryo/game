package main

import (
	"game/config"
	g "game/infrastructure/grpc"
	_ "github.com/go-sql-driver/mysql"
)



func main() {
	if err := g.CreateServer(); err != nil {
		config.Errorlog.Fatalf("failed to serve: %v", err)
	}
}
