package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"TP/configs"
	"TP/schema"
)

func convTSCurrency(c *gin.Context, tp *schema.TokenPrice, currency schema.Currency, level schema.TimeSeriesLevel) error {
	if res := configs.MongoPriceCol.FindOne(c, bson.M{"currency": currency, "level": level}); res.Err() != nil {
		return res.Err()
	} else {
		cp := schema.CurrencyPrice{}
		if err := res.Decode(&cp); err != nil {
			return err
		}
		if len(cp.TS) != len(tp.TS) {
			return fmt.Errorf("currency prices and token prices should have same no. of elements")
		}
		for i := 0; i < len(cp.TS); i++ {
			p := tp.TS[i].Price
			p *= cp.TS[i].Price
			tp.TS[i].Price = p
		}
	}
	return nil
}

func GetTokenPriceTS(c *gin.Context) {
	tokenId := c.Query("tokenId")
	if len(tokenId) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, -1)
		return
	}
	convRial := false
	currency := c.Query("currency")
	if len(currency) > 0 && currency != string(schema.Rial) {
		c.IndentedJSON(http.StatusUnprocessableEntity, fmt.Sprintf("Currency not yet supported! only %s", schema.Rial))
		return
	} else if currency == string(schema.Rial) {
		convRial = true
	}
	level := c.Query("level")
	if res := configs.MongoPriceCol.FindOne(c, bson.M{"token_id": tokenId, "level": level}); res.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, res.Err())
	} else {
		tp := schema.TokenPrice{}
		err := res.Decode(&tp)
		if convRial {
			if err := convTSCurrency(c, &tp, schema.Currency(currency), tp.Level); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(http.StatusOK, &tp)
	}
}

func GetTSLevels(c *gin.Context) {
	r := []schema.TimeSeriesLevel{schema.Hourly, schema.Daily, schema.Weekly, schema.Monthly, schema.Yearly}
	c.IndentedJSON(http.StatusOK, &r)
}
