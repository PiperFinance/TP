package main

import (
	"TP/bg"
	"TP/configs"
	"TP/views"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	fmt.Println("BOOT: Loading config ...")
	configs.LoadConfig()
	fmt.Println("BOOT: Loading Logger ...")
	configs.LoadLogger()
	fmt.Println("BOOT: Loading Cache ...")
	configs.LoadCache()
	fmt.Println("BOOT: Loading Mongo ...")
	configs.LoadMongo()
	fmt.Println("BOOT: Loading Tokens ...")
	configs.LoadTokens()
	fmt.Println("BOOT: Loading Wrapped - Tokens ...")
	configs.LoadWrapped()
	fmt.Println("BOOT: Loading Coin Market Cap ...")
	configs.LoadCMC()
}

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.GET("/currencies", views.GetCurrencies)
	router.GET("/stats", views.ServiceStats)
	router.GET("/", views.GetTokenPrice)
	router.POST("/", views.GetTokenPriceMulti)
	router.POST("/import", views.ImportHistory)
	router.GET("/all", views.GetAllTokensPrice)
	router.GET("/all/ids", views.GetAllTokenIds)

	// router.GET("/range", GetTokenPriceTS)
	router.GET("/ts/", views.GetTokenPriceTS)
	router.GET("/ts/flat", views.GetTokenPriceTSFlat)
	router.GET("/ts/levels", views.GetTokenPriceTS)

	bg.LoadBg()

	if err := router.Run(configs.Config.ApiUrl); err != nil {
		log.Fatal(err)
	}
}
