package configs

import "TP/schema"

var SupportedLevels = []schema.TimeSeriesLevel{
	schema.Hourly, schema.Daily, schema.Monthly, schema.Weekly, schema.Yearly,
}
