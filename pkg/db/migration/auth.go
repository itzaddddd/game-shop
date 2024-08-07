package migration

import (
	"context"
	"log"

	"github.com/itzaddddd/game-shop/config"
	"github.com/itzaddddd/game-shop/pkg/db"
	"github.com/itzaddddd/game-shop/service/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func authDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return db.DBConn(pctx, cfg).Database("auth_db")
}

func AuthMigrate(pctx context.Context, cfg *config.Config) {
	db := authDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	colAuth := db.Collection("auths")

	indexs, _ := colAuth.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
		{Keys: bson.D{{Key: "refresh_token", Value: 1}}},
	})

	for _, index := range indexs {
		log.Printf("index: %s\n", index)
	}

	colRole := db.Collection("roles")
	indexs, _ = colRole.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "code", Value: 1}}},
	})

	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	documents := func() (documents []interface{}) {
		roleDocs := []auth.Role{
			{
				Title: "player",
				Code:  0,
			},
			{
				Title: "admin",
				Code:  1,
			},
		}

		for _, role := range roleDocs {
			documents = append(documents, role)
		}

		return

	}()

	results, err := colRole.InsertMany(pctx, documents)

	if err != nil {
		panic(err)
	}

	log.Println("Migrate auth completed: ", results)

}
