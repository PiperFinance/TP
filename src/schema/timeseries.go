package schema

import "time"

type (
	TimeSeriesLevel string
	Currency        string
)

const (
	Rial  Currency = "IRR"
	USD   Currency = "USD"
	Pound Currency = "Pound"
	Yen   Currency = "Yen"
)

const (
	Hourly  TimeSeriesLevel = "H"
	Daily   TimeSeriesLevel = "D"
	Weekly  TimeSeriesLevel = "W"
	Monthly TimeSeriesLevel = "M"
	Yearly  TimeSeriesLevel = "Y"
)

type TimeSeries struct {
	Timestamp int64   `bson:"t" binding:"required" json:"t"`
	Price     float64 `bson:"p"  binding:"required" json:"p"`
}

func (ts *TimeSeries) GetTime() time.Time {
	return time.Unix(ts.Timestamp, 0)
}

type TokenPrice struct {
	TokenId    TokenId         `bson:"token_id" json:"token_id"`
	Level      TimeSeriesLevel `bson:"level" json:"level"`
	Detail     TokenDet        `json:"detail" bson:"detail"`
	LastUpdate time.Time       `bson:"last_update" json:"last_update"`
	TS         []TimeSeries    `json:"ts" bson:"ts"`
}

type CurrencyPrice struct {
	Currency   Currency        `bson:"currency" json:"currency"`
	Level      TimeSeriesLevel `bson:"level" json:"level"`
	Detail     TokenDet        `json:"detail" bson:"detail"`
	LastUpdate time.Time       `bson:"last_update" json:"last_update"`
	TS         []TimeSeries    `json:"ts" bson:"ts"`
}
