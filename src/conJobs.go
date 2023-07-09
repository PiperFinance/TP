package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eko/gocache/v3/store"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"TP/configs"
	"TP/nobitex"
	"TP/schema"
	"TP/utils"
)

func NobitexUSDTPrice() {
	if price, err := nobitex.RialPrice(); err != nil {
		log.Errorf("CMCErr: %+v", err)
	} else {
		cacheKey := "NBT:USDT-Rial"
		_ = configs.TokenPriceCache.Set(context.Background(), cacheKey, price, store.WithExpiration(NobitexApiTTL))
	}
}

func CMCPrices() {
	now := time.Now()
	if _, prices, err := cmcUSDClient.AllPrices(); err != nil {
		log.Errorf("CMCErr: %+v", err)
	} else {
		for tokenId, price := range prices {
			_ = configs.TokenPriceCache.Set(
				context.Background(),
				fmt.Sprintf("CCT:%s", tokenId),
				price,
				store.WithExpiration(CMCTPMultiTTL))
			token, ok := configs.AllTokens[tokenId]
			if ok {
				_ = configs.TokenPriceCache.Set(
					context.Background(),
					fmt.Sprintf("CCT:GID:%s", token.Detail.CoingeckoId),
					price,
					store.WithExpiration(CMCTPMultiTTL))
			}
			configs.AllIdsMutex.Lock()
			configs.AllIds[tokenId] = now
			configs.AllIdsMutex.Lock()
		}
	}
}

func CGPrices() {
	ids := make([]string, 1)
	for _, token := range configs.AllChainsTokens() {
		if len(token.Detail.CoingeckoId) > 0 {
			ids = append(ids, token.Detail.CoingeckoId)
		}
	}
	chunks := utils.Chunks(ids, CGPriceUpdaterChunkSize)
	counter := 0
	for _, chunk := range chunks {
		// TODO - maybe multi referenced price
		// TODO - Add proxy here ...
		// TODO - add token Ids
		res, err := cg.SimplePrice(chunk, []string{"usd"})
		if err != nil {
			log.Error(err)
		} else {
			counter++
			for geckoId, currencies := range *res {
				for _, tokenId := range configs.GeckoIdToTokenId(geckoId) {
					_ = configs.TokenPriceCache.Set(
						context.Background(),
						fmt.Sprintf("CGT:%s", tokenId),
						float64(currencies["usd"]),
						store.WithExpiration(CGTPMultiTTL))
					configs.AllIdsMutex.Lock()
					configs.AllIds[tokenId] = time.Now()
					configs.AllIdsMutex.Unlock()

				}
				_ = configs.TokenPriceCache.Set(
					context.Background(),
					fmt.Sprintf("CGT:ID:%s", geckoId),
					float64(currencies["usd"]),
					store.WithExpiration(CGTPMultiTTL))
			}
			time.Sleep(CGTPMultiCallDelay)
			log.Infof("--successfully fetched %d", counter)
		}
	}
}

type aggCount struct {
	ID    primitive.ObjectID `bson:_id`
	count int                `bson:"count"`
}

func initToken(ctx context.Context, tokenId schema.TokenId, ts schema.TimeSeries, lastUpdate time.Time) error {
	token, ok := configs.AllTokens[tokenId]
	if !ok {
		// TODO - show warning here
		return nil
	}
	for _, level := range configs.SupportedLevels {
		tp := schema.TokenPrice{
			TokenId:    tokenId,
			Level:      level,
			Detail:     token.Detail,
			LastUpdate: lastUpdate,
			TS:         []schema.TimeSeries{ts},
		}
		if _, err := configs.MongoPriceCol.InsertOne(ctx, &tp); err != nil {
			return err
		}
	}
	configs.TSIds[tokenId] = true
	return nil
}

func checkMaxSize(ctx context.Context, level schema.TimeSeriesLevel, maxArrSize int) error {
	p := bson.A{
		bson.D{{
			Key:   "$match",
			Value: bson.D{{Key: "level", Value: level}},
		}},
		bson.D{{
			Key:   "$project",
			Value: bson.D{{Key: "count", Value: bson.M{"$size": "$ts"}}, {Key: "_id", Value: "$_id"}},
		}},
	}
	cur, err := configs.MongoPriceCol.Aggregate(ctx, p)
	if err != nil {
		log.Error(err)
		return err
	}
	manyIds := make([]primitive.ObjectID, 0)
	aggArr := []aggCount{}
	if err := cur.All(ctx, &aggArr); err != nil {
		return err
	}
	for _, agg := range aggArr {
		if err := cur.Decode(&agg); err != nil {
			return err
		}
		if agg.count >= maxArrSize {
			manyIds = append(manyIds, agg.ID)
		}
	}
	if len(manyIds) > 0 {
		if _, err := configs.MongoPriceCol.UpdateMany(
			ctx,
			bson.M{"_id": bson.M{"$in": &manyIds}},
			bson.M{"$pop": bson.M{"ts": 1}},
		); err != nil {
			return err
		}
	}
	return nil
}

func updateCurrency(ctx context.Context, level schema.TimeSeriesLevel, now time.Time, ts schema.TimeSeries) error {
	if count, err := configs.MongoPriceCol.CountDocuments(ctx, bson.M{"level": level, "currency": schema.Rial}); err != nil && err != mongo.ErrNoDocuments {
		return err
	} else if count == 0 {
		cp := schema.CurrencyPrice{
			Currency:   schema.Rial,
			Level:      level,
			LastUpdate: now,
			TS:         []schema.TimeSeries{ts},
		}
		if _, err := configs.MongoPriceCol.InsertOne(ctx, &cp); err != nil {
			return err
		}
	} else {
		if _, err := configs.MongoPriceCol.UpdateOne(
			ctx,
			bson.M{"currency": schema.Rial, "level": level},
			bson.M{"$push": bson.M{"ts": &ts}},
		); err != nil {
			log.Error(err)
		}
	}
	return nil
}

func updateTokens(ctx context.Context, level schema.TimeSeriesLevel, now time.Time) error {
	for tokenId, lastUpdate := range configs.AllIds {
		price := getPriceRaw(ctx, "USD", tokenId)
		ts := schema.TimeSeries{Timestamp: now, Price: price}
		if count, err := configs.MongoPriceCol.CountDocuments(ctx, bson.M{"level": level, "token_id": tokenId}); err != nil && err != mongo.ErrNoDocuments {
			return err
		} else if count == 0 {
			if err := initToken(ctx, tokenId, ts, lastUpdate); err != nil {
				log.Error(err)
				return err
			}
		} else {
			if _, err := configs.MongoPriceCol.UpdateOne(
				ctx,
				bson.M{"token_id": tokenId, "level": level},
				bson.M{"$push": bson.M{"ts": &ts}},
			); err != nil {
				log.Error(err)
			}
		}

	}
	return nil
}

func TSCapture(level schema.TimeSeriesLevel, maxArrSize int) error {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Config.CaptureTimeSeriesTaskTimeout)
	defer cancel()
	now := time.Now()
	cts := schema.TimeSeries{Timestamp: now, Price: nobitex.LastIRRPrice(ctx)}
	if err := updateCurrency(ctx, level, now, cts); err != nil {
		return err
	}
	if err := updateTokens(ctx, level, now); err != nil {
		return err
	}
	if err := checkMaxSize(ctx, level, maxArrSize); err != nil {
		return err
	}
	return nil
}

func TSSnapshotHourly() {
	if err := TSCapture(schema.Hourly, (60)/1); err != nil {
		log.Error(err)
	}
}

func TSSnapshotDaily() {
	if err := TSCapture(schema.Daily, (60*24)/5); err != nil {
		log.Error(err)
	}
}

func TSSnapshotWeekly() {
	if err := TSCapture(schema.Weekly, (60*24*7)/5); err != nil {
		log.Error(err)
	}
}

func TSSnapshotMonthly() {
	if err := TSCapture(schema.Monthly, (60*24*30)/30); err != nil {
		log.Error(err)
	}
}

func TSSnapshotYearly() {
	if err := TSCapture(schema.Yearly, (60*24*366)/(6*60)); err != nil {
		log.Error(err)
	}
}
