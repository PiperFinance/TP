package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"TP/configs"
	"TP/nobitex"
	"TP/schema"
)

func getPrice(c *gin.Context, tokenId schema.TokenId) float64 {
	multiplier := float64(0)
	currencyQ := c.Query("currency")
	switch currencyQ {
	case "":
		multiplier = 1
	case "USD":
		multiplier = 1
	case "IRR":
		multiplier = nobitex.LastIRRPrice(c)
	}
	// NOTE - to change usd to other currencies, by default is USD
	var tokenPrice float64
	geckoPrice, _ := strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CGT:%s", tokenId)).Val(), 64)
	cmcPrice, _ := strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CCT:%s", tokenId)).Val(), 64)
	if cmcPrice == 0 && geckoPrice == 0 {
		// NOTE - if 0xeeeee... not found tries it's wrapped values !
		tokenId, ok := configs.WrappedTokensMap[tokenId]
		if !ok {
			return 0
		}
		geckoPrice, _ = strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CGT:%s", tokenId)).Val(), 64)
		cmcPrice, _ = strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CCT:%s", tokenId)).Val(), 64)
	}
	if cmcPrice > 0 && geckoPrice > 0 {
		tokenPrice = (cmcPrice + geckoPrice) / 2
	} else if cmcPrice > 0 {
		tokenPrice = cmcPrice
	} else if geckoPrice > 0 {
		tokenPrice = geckoPrice
	}

	return tokenPrice * multiplier
}

func GetTokenPriceMulti(c *gin.Context) {
	ids := make([]schema.TokenId, 0)
	res := make(map[schema.TokenId]float64)
	err := c.BindJSON(&ids)
	if err != nil {
		log.Error(err)
		return
	}

	for _, tokenId := range ids {
		res[tokenId] = getPrice(c, tokenId)
	}

	c.JSON(http.StatusOK, res)
}

func GetTokenPrice(c *gin.Context) {
	tokenId := c.Query("tokenId")
	if len(tokenId) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, -1)
		return
	}
	price := getPrice(c, schema.TokenId(tokenId))
	if price == 0 {
		c.IndentedJSON(http.StatusTooEarly, 0)
	} else {
		c.IndentedJSON(http.StatusOK, price)
	}
}

type Stats struct {
	CoinGecko        int           `json:"cg"`
	CoinGeckoTTL     time.Duration `json:"cgTTL"`
	CoinMarketCap    int           `json:"cmc"`
	CoinMarketCapTTL time.Duration `json:"cmcTTL"`
	Nobitex          float64       `json:"rial"`
	NobitexTTL       time.Duration `json:"rialTTL"`
}

func ServiceStats(c *gin.Context) {
	r := Stats{}
	if cmd := configs.RedisClient.Keys(c, "CGT:*"); cmd.Err() == nil {
		r.CoinGecko = len(cmd.Val())
		if r.CoinGecko > 0 {
			if cmd := configs.RedisClient.TTL(c, cmd.Val()[0]); cmd.Err() == nil {
				r.CoinGeckoTTL = cmd.Val().Abs()
			}
		}
	}
	if cmd := configs.RedisClient.Keys(c, "CCT:*"); cmd.Err() == nil {
		r.CoinMarketCap = len(cmd.Val())
		if r.CoinMarketCap > 0 {
			if cmd := configs.RedisClient.TTL(c, cmd.Val()[0]); cmd.Err() == nil {
				r.CoinMarketCapTTL = cmd.Val().Abs()
			}
		}
	}
	if cmd := configs.RedisClient.Get(c, "NBT:USDT-Rial"); cmd.Err() == nil {
		r.Nobitex, _ = strconv.ParseFloat(cmd.Val(), 64)
		if r.Nobitex > 0 {
			if cmd := configs.RedisClient.TTL(c, "NBT:USDT-Rial"); cmd.Err() == nil {
				r.NobitexTTL = cmd.Val().Abs()
			}
		}
	}
	c.IndentedJSON(200, r)
}
