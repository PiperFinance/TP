package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eko/gocache/v3/store"
	log "github.com/sirupsen/logrus"

	"TP/configs"
	"TP/nobitex"
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
	if _, prices, err := cmcUSDClient.AllPrices(); err != nil {
		log.Errorf("CMCErr: %+v", err)
	} else {
		for tokenId, price := range prices {
			cacheKey := fmt.Sprintf("CCT:%s", tokenId)
			_ = configs.TokenPriceCache.Set(context.Background(), cacheKey, price, store.WithExpiration(CMCTPMultiTTL))
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
		res, err := cg.SimplePrice(chunk, []string{"usd"})
		if err != nil {
			log.Error(err)
		} else {
			counter++
			for geckoId, currencies := range *res {
				cacheKey := fmt.Sprintf("CGT:%s", configs.GeckoIdToTokenId(geckoId))
				_ = configs.TokenPriceCache.Set(context.Background(), cacheKey, float64(currencies["usd"]), store.WithExpiration(CGTPMultiTTL))
			}
			time.Sleep(CGTPMultiCallDelay)
			log.Infof("--successfully fetched %d", counter)
		}
	}
}
