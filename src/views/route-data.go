package views

import (
	"TP/configs"
	"TP/schema"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type History struct {
	TokenId schema.TokenId         `json:"tokenId"`
	Data    HistoryData            `json:"data"`
	Level   schema.TimeSeriesLevel `json:"level"`
}
type HistoryPrice struct {
	Timestamp int64
	Price     float64
}
type HistoryData struct {
	CoinUSD []HistoryPrice `json:"price"`  // Price is Dollar [ (ts :int, pr: float ) , () , ...  ]
	USDIRR  []HistoryPrice `json:"dollar"` // Price of Dollar into Rial [ (ts :int, pr: float ) , () , ...  ]
}

func (t *HistoryPrice) UnmarshalJSON(b []byte) error {
	a := []interface{}{&t.Timestamp, &t.Price}
	return json.Unmarshal(b, &a)
}

// func (t *myType) UnmarshalJSON(b []byte) error {
//    a := []interface{}{&t.count, &t.name, &t.relation}
//    return json.Unmarshal(b, &a)
// }

func (h *History) commitDB(c *gin.Context) error {
	token, ok := configs.AllTokens[h.TokenId]
	if !ok {
		return fmt.Errorf("token id %s not found", h.TokenId)
	}
	coinTs := make([]schema.TimeSeries, 0)
	tp := schema.TokenPrice{
		TokenId:    h.TokenId,
		Level:      h.Level,
		Detail:     token.Detail,
		LastUpdate: time.Now(),
		TS:         coinTs,
	}
	// TODO - Maybe make USDIRR sorted and stop at middle to improve speed
	// TOOD - add usd prices to mongo as well
	for _, coinPrice := range h.Data.CoinUSD {
		ok := false
		for _, usdPrice := range h.Data.USDIRR {
			if usdPrice.Timestamp == coinPrice.Timestamp {
				ok = true
				break
			}
		}
		if ok {
			coinTs = append(coinTs, schema.TimeSeries{
				Timestamp: coinPrice.Timestamp,
				Price:     coinPrice.Price,
			})
		}
	}
	tp.TS = coinTs
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"token_id": h.TokenId, "level": h.Level}
	update := bson.D{{Key: "$set", Value: &tp}}
	_, err := configs.MongoPriceCol.UpdateOne(c, filter, update, opts)
	return err
}

func ImportHistory(c *gin.Context) {
	history := History{}
	err := c.BindJSON(&history)
	if err != nil {
		c.AbortWithError(422, err)
	} else {
		history.commitDB(c)
		c.IndentedJSON(200, &history)
	}
}
