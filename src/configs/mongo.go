package configs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// LogColName Collection name for transfers events
	LogColName          = "Logs"
	BlockColName        = "Blocks"
	TokenColName        = "Tokens"
	ParsedLogColName    = "ParsedLogs"
	UserBalColName      = "UsersBalance"
	BannedUsersColName  = "BannedUsers"
	TransfersColName    = "Transfers"
	TokenVolumeColName  = "TokenVolume"
	TokenUserMapColName = "TokenUserMap"
	UserTokenMapColName = "UserTokenMap"
	QueueErrorsColName  = "QErr"
	TokenPriceDB        = "TP"
	PriceCol            = "TokenPriceTS"
	BlockScannerDB      = "BS_Main"
	AggregatedUsers     = "Users"
)

var (
	mongoCl       *mongo.Client
	MongoPriceCol *mongo.Collection
)

func LoadMongo() {
	time.Sleep(Config.MongoSlowLoading)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI(Config.MongoUrl.String())
	opts.MaxPoolSize = &Config.MongoMaxPoolSize
	var err error
	mongoCl, err = mongo.Connect(ctx, opts)
	if err != nil {
		Logger.Panicf("Mongo: %s", err)
	}

	err = mongoCl.Ping(ctx, nil)
	if err != nil {
		Logger.Panicf("Mongo: %s", err)
	}

	MongoPriceCol = mongoCl.Database(TokenPriceDB).Collection(PriceCol)

	MongoPriceCol.Indexes().CreateOne(
		ctx, mongo.IndexModel{
			Keys: bson.D{{Key: "token_id", Value: 1}, {Key: "level", Value: 1}},
		})

	MongoPriceCol.Indexes().CreateOne(
		ctx, mongo.IndexModel{
			Keys: bson.D{{Key: "currency", Value: 1}, {Key: "level", Value: 1}},
		})
	MongoPriceCol.Indexes().CreateOne(
		ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "token_id", Value: 1}, {Key: "level", Value: 1}, {Key: "currency", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
}
