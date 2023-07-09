package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron"
	coingecko "github.com/superoo7/go-gecko/v3"

	"TP/configs"
	simplecmc "TP/simpleCMC"
)

var (
	cg           *coingecko.Client
	cmcUSDClient simplecmc.Client
)

const (
	NobitexUSDTPriceUpdaterCron = "0 */1 * * * *"
	NobitexApiTTL               = 1 * 5 * time.Minute

	TSHourlyCron  = "59 */1 * * * *"  // NOTE - Every 1 minutes -> ( 60  ) / 1  = 60
	TSDailyCron   = "59 */5 * * * *"  // NOTE - Every 5 minutes -> ( 60  * 24 ) / 5  = 288
	TSWeeklyCron  = "59 */15 * * * *" // NOTE - Every 15 minutes -> ( 60 * 24 * 7 ) / 5  = 2016
	TSMonthlyCron = "59 */29 * * * *" // NOTE - Every 30 minutes -> ( 60 * 24 * 30 ) / 30 = 1440
	TSYearlyCron  = "59 59 */5 * * *" // NOTE - Every Five minutes -> ( 60 * 24 * 365 ) / ( 6 * 60 ) = 1460

	CMCPriceUpdaterCron = "0 */5 * * * *"
	CMCTPMultiTTL       = 5 * 5 * time.Minute

	// NOTE Coin Gecko Public api
	// total of 9000 tokens with 10 req/minute
	// call count  ::: 9000 / 400 = 22 -> 10 - 10 - 2 => 3 minutes
	// calls are limited cause GET -> query params
	CGPriceUpdaterCron      = "0 */45 * * * *"
	CGPriceUpdaterChunkSize = 400
	CGCallTimeout           = 5 * time.Second
	CGTPMultiCallDelay      = 1 * time.Second
	CGTPMultiTTL            = 45 * 2 * time.Minute
)

func init() {
	configs.LoadConfig()
	configs.LoadLogger()
	configs.LoadCache()
	configs.LoadMongo()
	configs.LoadTokens()
	configs.LoadWrapped()
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
	router.GET("/currencies", GetCurrencies)
	router.GET("/stats", ServiceStats)
	router.GET("/", GetTokenPrice)
	router.POST("/", GetTokenPriceMulti)
	router.GET("/all", GetAllTokensPrice)
	router.GET("/all/ids", GetAllTokenIds)

	// router.GET("/range", GetTokenPriceTS)
	router.GET("/ts/", GetTokenPriceTS)
	router.GET("/ts/levels", GetTokenPriceTS)

	cr := cron.New()
	if err := cr.AddFunc(NobitexUSDTPriceUpdaterCron, NobitexUSDTPrice); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(CMCPriceUpdaterCron, CMCPrices); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(CGPriceUpdaterCron, CGPrices); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(TSHourlyCron, TSSnapshotHourly); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(TSDailyCron, TSSnapshotDaily); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(TSWeeklyCron, TSSnapshotWeekly); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(TSMonthlyCron, TSSnapshotMonthly); err != nil {
		log.Panic(err)
	}
	if err := cr.AddFunc(TSYearlyCron, TSSnapshotYearly); err != nil {
		log.Panic(err)
	}
	cr.Start()

	if err := router.Run(configs.Config.ApiUrl); err != nil {
		log.Fatal(err)
	}
}
