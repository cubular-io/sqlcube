package sqlc

type InsertDailyParams struct {
	TickerID int32   `json:"ticker_id"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Time     int64   `json:"time"`
}

type InsertHourlyParams struct {
	TickerID int32   `json:"ticker_id"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Time     int64   `json:"time"`
}

type InsertMinuteParams struct {
	TickerID int32   `json:"ticker_id"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Time     int64   `json:"time"`
}
