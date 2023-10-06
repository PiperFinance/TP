package views

import (
	"TP/configs"
	"TP/nobitex"
	"TP/schema"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetPrice(c *gin.Context, tokenId schema.TokenId) float64 {
	return GetPriceRaw(c, c.Query("currency"), tokenId)
}

func GetPriceRaw(c context.Context, currencyQ string, tokenId schema.TokenId) float64 {
	multiplier := float64(0)
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
	// NOTE - in case no results were found again
	if cmcPrice == 0 && geckoPrice == 0 {
		token, ok := configs.AllTokens[tokenId]
		if ok {
			cmcPrice, _ = strconv.ParseFloat(
				configs.RedisClient.Get(
					c,
					fmt.Sprintf("CCT:GID:%s", token.Detail.CoingeckoId)).Val(), 64)
		}
	}
	if cmcPrice == 0 && geckoPrice == 0 {
		// NOTE - if 0xeeeee... not found tries it's wrapped values !
		tokenId, ok := configs.WrappedTokensMap[tokenId]
		if ok {
			geckoPrice, _ = strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CGT:%s", tokenId)).Val(), 64)
			cmcPrice, _ = strconv.ParseFloat(configs.RedisClient.Get(c, fmt.Sprintf("CCT:%s", tokenId)).Val(), 64)
		}
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
		res[tokenId] = GetPrice(c, tokenId)
	}

	c.JSON(http.StatusOK, res)
}

func GetAllTokensPrice(c *gin.Context) {
	res := make(map[schema.TokenId]float64)
	includeZeros, _ := strconv.ParseBool(c.Query("includeZeros"))
	configs.AllIdsMutex.RLock()
	for tokenId := range configs.AllIds {
		price := GetPrice(c, tokenId)
		if !includeZeros && price == 0 {
			continue
		}
		res[tokenId] = price
	}
	configs.AllIdsMutex.RUnlock()
	c.JSON(http.StatusOK, res)
}

func GetAllTokenIds(c *gin.Context) {
	configs.AllIdsMutex.RLock()
	c.JSON(http.StatusOK, configs.AllIds)
	configs.AllIdsMutex.RUnlock()
}

func GetCurrencies(c *gin.Context) {
	r := []schema.Currency{schema.Rial, schema.USD}
	c.IndentedJSON(http.StatusOK, &r)
}

func GetTokenPrice(c *gin.Context) {
	tokenId := c.Query("tokenId")
	if len(tokenId) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, -1)
		return
	}
	price := GetPrice(c, schema.TokenId(tokenId))
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
