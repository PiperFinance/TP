package main

import (
	"TP/configs"
	"TP/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"strconv"
	"time"
)

var cg *coingecko.Client

func init() {
	fmt.Println("HI")
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg = coingecko.NewClient(httpClient)

}

func main() {
	fmt.Println("StartingApp")
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/", GetTokenPrice)

	err := router.Run(fmt.Sprintf("0.0.0.0:%s", configs.GetAppPort()))
	if err != nil {
		log.Fatal(err)
	}
}

func GetTokenPrice(c *gin.Context) {
	//configs.GetToken(c.Query(''))

	tokenId, err := strconv.ParseInt(c.Query("tokenId"), 10, 32)
	if err != nil {
		log.Error(err)
	}
	token := configs.GetToken(schema.TokenId(tokenId))
	coin, err := cg.CoinsID(
		token.Detail.CoingeckoId, false, false, true, false, false, false)
	if err != nil {
		log.Error(err)
		return
	}
	c.IndentedJSON(http.StatusOK, coin.MarketData.CurrentPrice["usd"])

}
