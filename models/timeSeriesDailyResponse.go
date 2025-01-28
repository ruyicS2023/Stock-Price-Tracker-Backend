package models

// TimeSeriesDailyResponse represents the API response for daily stock data (TIME_SERIES_DAILY).
type TimeSeriesDailyResponse struct {
	MetaData        map[string]string    `json:"Meta Data"`
	TimeSeriesDaily map[string]DailyData `json:"Time Series (Daily)"`
}
