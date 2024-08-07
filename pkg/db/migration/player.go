package migration

import (
	"context"
	"log"

	"github.com/itzaddddd/game-shop/config"
	"github.com/itzaddddd/game-shop/pkg/db"
	"github.com/itzaddddd/game-shop/pkg/utils"
	"github.com/itzaddddd/game-shop/service/player"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func playerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return db.DBConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config) {
	playerDb := playerDbConn(pctx, cfg)
	defer playerDb.Client().Disconnect(pctx)

	playerTxCol := playerDb.Collection("player_transactions")

	indexs, _ := playerTxCol.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "_id", Value: 1}}},
		{Keys: bson.D{{Key: "player_id", Value: 1}}},
	})
	log.Println(indexs)

	playerCol := playerDb.Collection("players")

	documents := func() (documents []interface{}) {
		players := []*player.Player{
			{
				Email: "player001@email.com",
				Password: func() string {
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "player001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "player002@email.com",
				Password: func() string {
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "player002",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "player003@email.com",
				Password: func() string {
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "player003",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "admin001@email.com",
				Password: func() string {
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "admin001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		for _, player := range players {
			documents = append(documents, player)
		}

		return
	}()

	results, err := playerCol.InsertMany(pctx, documents)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player completed: ", results)

	playerTransactions := make([]interface{}, 0)
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, player.PlayerTransaction{
			PlayerId:  "player:" + p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime(),
		})
	}

	results, err = playerTxCol.InsertMany(pctx, playerTransactions)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player_transactions completed: ", results)

	playerTxQueueCol := playerDb.Collection("player_transactions_queue")

	result, err := playerTxQueueCol.InsertOne(pctx, bson.M{"offset": -1})

	if err != nil {
		panic(err)
	}
	log.Println("Migrate player_transactions_queue completed: ", result)

}
