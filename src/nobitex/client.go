package nobitex

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"

	"TP/configs"
)

const apiUrl = "https://api.nobitex.ir/market/stats?srcCurrency=usdt&dstCurrency=rls"

func RialPrice() (float64, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	r := MarketPrice{}
	if err := json.Unmarshal(body, &r); err != nil {
		return 0, err
	}
	buy, err := strconv.ParseFloat(r.Stats.UsdtRls.BestBuy, 64)
	if err != nil {
		return 0, err
	}
	sell, err := strconv.ParseFloat(r.Stats.UsdtRls.BestSell, 64)
	if err != nil {
		return 0, err
	}
	return (buy + sell) / 2, nil
}

func LastIRRPrice(c context.Context) float64 {
	cacheKey := "NBT:USDT-Rial"
	v := configs.RedisClient.Get(c, cacheKey).Val()
	if price, err := strconv.ParseFloat(v, 64); err == redis.Nil {
		logrus.Error(err)
		new, err := RialPrice()
		if err != nil {
			logrus.Error(err)
		}
		return new
	} else {
		return (price)
	}
}
