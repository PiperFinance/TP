package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	coingecko "github.com/superoo7/go-gecko/v3"

	"TP/configs"
	simplecmc "TP/simpleCMC"
)

var (
	cg           *coingecko.Client
	cmcUSDClient simplecmc.Client
)

const (
	NobitexUSDTPriceUpdaterCron = "*/1 * * * *"
	NobitexApiTTL               = 1 * 5 * time.Minute

	CMCPriceUpdaterCron = "*/10 * * * *"
	CMCTPMultiTTL       = 10 * 5 * time.Minute

	// NOTE Coin Gecko Public api
	// total of 9000 tokens with 10 req/minute
	// call count  ::: 9000 / 400 = 22 -> 10 - 10 - 2 => 3 minutes
	// calls are limited cause GET -> query params
	CGPriceUpdaterCron      = "*/45 * * * *"
	CGPriceUpdaterChunkSize = 400
	CGCallTimeout           = 5 * time.Second
	CGTPMultiCallDelay      = 1 * time.Second
	CGTPMultiTTL            = 45 * 2 * time.Minute
)

func init() {
	configs.LoadConfig()
	configs.LoadLogger()
	configs.LoadCache()
	configs.LoadCMC()
	cmcUSDClient = simplecmc.Client{
		ApiKeys:  configs.Config.CMCApiKeys,
		Currency: "USD",
	}
	httpClient := &http.Client{
		Timeout: CGCallTimeout,
	}
	cg = coingecko.NewClient(httpClient)
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/stats", ServiceStats)
	router.GET("/", GetTokenPrice)
	router.POST("/", GetTokenPriceMulti)

	cr := cron.New()
	if _, err := cr.AddFunc(NobitexUSDTPriceUpdaterCron, NobitexUSDTPrice); err != nil {
		log.Panic(err)
	}
	if _, err := cr.AddFunc(CMCPriceUpdaterCron, CMCPrices); err != nil {
		log.Panic(err)
	}
	if _, err := cr.AddFunc(CGPriceUpdaterCron, CGPrices); err != nil {
		log.Panic(err)
	}
	cr.Start()

	if err := router.Run(configs.Config.ApiUrl); err != nil {
		log.Fatal(err)
	}
}
