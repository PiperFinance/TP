package main

import (
	"TP/configs"
	"TP/schema"
	"TP/utils"
	"context"
	"fmt"
	"github.com/eko/gocache/v3/store"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"strconv"
	"time"
)

var cg *coingecko.Client

func init() {
	httpClient := &http.Client{
		Timeout: time.Second * 8,
	}
	cg = coingecko.NewClient(httpClient)

}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/", GetTokenPrice)
	router.POST("/", GetTokenPriceMulti)

	cr := cron.New()
	_, err := cr.AddFunc("*/5 * * * *", func() {
		ids := make([]string, 1)
		for _, token := range configs.AllChainsTokens() {
			if len(token.Detail.CoingeckoId) > 0 {
				ids = append(ids, token.Detail.CoingeckoId)
			}
		}
		chunks := utils.Chunks(ids, 450)
		counter := 0
		for _, chunk := range chunks {
			res, err := cg.SimplePrice(chunk, []string{"usd"})
			if err != nil {
				log.Error(err)
			} else {
				counter++
				for geckoId, currencies := range *res {
					cacheKey := fmt.Sprintf("CT:%d", configs.GeckoIdToTokenId(geckoId))
					_ = configs.TokenPriceCache.Set(context.Background(), cacheKey, float64(currencies["usd"]), store.WithExpiration(5*time.Minute))
				}
				time.Sleep(15 * time.Second)
				log.Infof("--succefully fetched %d", counter)
			}
		}

	})
	cr.Start()

	err = router.Run(fmt.Sprintf("0.0.0.0:%s", configs.GetAppPort()))
	if err != nil {
		log.Fatal(err)
	}
}

func GetTokenPriceMulti(c *gin.Context) {
	ids := make([]schema.TokenId, 1)
	res := make(map[schema.TokenId]float64)
	err := c.BindJSON(&ids)
	if err != nil {
		log.Error(err)
		return
	}
	for _, tokenId := range ids {
		cacheKey := fmt.Sprintf("CT:%d", tokenId)
		tokenPrice, _ := configs.TokenPriceCache.Get(context.Background(), cacheKey)
		res[tokenId] = tokenPrice
	}
	c.JSON(http.StatusOK, res)

}

func GetTokenPrice(c *gin.Context) {

	_tokenId := c.Query("tokenId")
	if len(_tokenId) == 0 {
		return
	}

	tokenId, err := strconv.ParseInt(_tokenId, 10, 32)
	cacheKey := fmt.Sprintf("CT:%d", tokenId)

	if err != nil {
		log.Error(err)
	}
	token := configs.GetToken(schema.TokenId(tokenId))

	tokenPrice, _ := configs.TokenPriceCache.Get(context.Background(), cacheKey)
	if tokenPrice == 0 {
		coin, err := cg.CoinsID(
			token.Detail.CoingeckoId, false, false, true, false, false, false)
		if err != nil {
			log.Error(err)
		}
		if err != nil || coin != nil {
			_ = configs.TokenPriceCache.Delete(context.Background(), cacheKey)
		} else {
			tokenPrice, _ = coin.MarketData.CurrentPrice["usd"]
			_ = configs.TokenPriceCache.Set(context.Background(), cacheKey, tokenPrice, store.WithExpiration(30*time.Second))
		}
	}
	log.Infof("TID: %s  R: %s", tokenId, tokenPrice)
	c.IndentedJSON(http.StatusOK, tokenPrice)

}
