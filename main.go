package main

import (
	"context"
	"log"
	"os"

	"github.com/itzaddddd/game-shop/config"
	"github.com/itzaddddd/game-shop/pkg/db"
	"github.com/itzaddddd/game-shop/server"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}

		return os.Args[1]
	}())

	db := db.DBConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
