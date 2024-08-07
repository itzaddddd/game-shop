package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/itzaddddd/game-shop/pkg/utils"
	"github.com/itzaddddd/game-shop/service/player"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionPlayers   = "players"
	CollectionPlayersTx = "player_transactions"
	Timeout             = 10 * time.Second
)

type (
	PlayerRepositoryService interface {
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTranscation(pctx context.Context, req *player.PlayerTransaction) (primitive.ObjectID, error)
		GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error)
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playerRepository{db: db}
}

func (r *playerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}

func (r *playerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, Timeout)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection(CollectionPlayers)

	player := new(player.Player)
	if err := col.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}

	return false
}

func (r *playerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, Timeout)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection(CollectionPlayers)

	res, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %s", err.Error())
		return primitive.NilObjectID, errors.New("insert one player failed")
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (r *playerRepository) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection(CollectionPlayers)

	result := new(player.PlayerProfileBson)

	if err := col.FindOne(pctx,
		bson.M{"_id": utils.ConvertToObjectId(playerId)},
		options.FindOne().SetProjection(bson.M{
			"_id":        1,
			"email":      1,
			"username":   1,
			"created_at": 1,
			"updated_at": 1,
		}),
	).
		Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("player profile not found")
	}

	return result, nil
}

func (r *playerRepository) InsertOnePlayerTranscation(pctx context.Context, req *player.PlayerTransaction) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection(CollectionPlayersTx)

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayerTranscation: %s", err.Error())
		return primitive.NilObjectID, errors.New("insert one player transcation failed")
	}
	log.Printf("Result: InsertOnePlayerTranscation: %v", result.InsertedID)

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *playerRepository) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	filter := bson.A{
		bson.D{{"$match", bson.D{{"player_id", playerId}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$player_id"},
					{"balance", bson.D{{"$sum", "$amount"}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"player_id", "$_id"},
					{"_id", 0},
					{"balance", 1},
				},
			},
		},
	}

	cursors, err := col.Aggregate(ctx, filter)
	if err != nil {
		log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
		return nil, errors.New("error: failed to get player saving account")
	}

	result := new(player.PlayerSavingAccount)
	for cursors.Next(ctx) {
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
			return nil, errors.New("error: failed to get player saving account")
		}
	}

	return result, nil
}