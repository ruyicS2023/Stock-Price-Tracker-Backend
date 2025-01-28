package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"stock-price-tracker/models"
)

const (
	baseURL = "https://www.alphavantage.co/query"
	apiKey  = "X5JYYRNVLPKN3MQX"
)

// AlphaVantageParams holds parameters for any Alpha Vantage endpoint.
type AlphaVantageParams struct {
	Function   string // e.g. "TIME_SERIES_DAILY", "GLOBAL_QUOTE", etc.
	Symbol     string
	Interval   string // for intraday, e.g., "5min"
	OutputSize string // "compact" or "full"
}

// FetchData is a generic helper that calls the Alpha Vantage API.
func FetchData(params AlphaVantageParams) ([]byte, error) {
	url := fmt.Sprintf("%s?function=%s&symbol=%s&apikey=%s", baseURL, params.Function, params.Symbol, apiKey)

	// Append optional parameters if needed
	if params.Interval != "" {
		url += "&interval=" + params.Interval
	}
	if params.OutputSize != "" {
		url += "&outputsize=" + params.OutputSize
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// FetchGlobalQuote fetches the current (real-time) stock data from Alpha Vantage.
func FetchGlobalQuote(symbol string) (models.GlobalQuoteResponse, error) {
	p := AlphaVantageParams{
		Function: "GLOBAL_QUOTE",
		Symbol:   symbol,
	}

	body, err := FetchData(p)
	if err != nil {
		return models.GlobalQuoteResponse{}, err
	}

	var quoteResp models.GlobalQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return models.GlobalQuoteResponse{}, err
	}

	return quoteResp, nil
}

// FetchTimeSeriesDaily fetches daily historical data (up to 100 days for "compact") from Alpha Vantage.
func FetchTimeSeriesDaily(symbol string) (models.TimeSeriesDailyResponse, error) {
	p := AlphaVantageParams{
		Function:   "TIME_SERIES_DAILY",
		Symbol:     symbol,
		OutputSize: "compact", // or "full" for 20+ years (Alpha Vantage limitations apply)
	}

	body, err := FetchData(p)
	if err != nil {
		return models.TimeSeriesDailyResponse{}, err
	}

	var dailyResp models.TimeSeriesDailyResponse
	if err := json.Unmarshal(body, &dailyResp); err != nil {
		return models.TimeSeriesDailyResponse{}, err
	}

	return dailyResp, nil
}

// FetchQuoteTrending fetches the latest price and volume information for a ticker.
func FetchQuoteTrending(symbol string) (models.QuoteTrendingResponse, error) {
	p := AlphaVantageParams{
		Function: "GLOBAL_QUOTE", // This function returns the latest price & volume
		Symbol:   symbol,
	}

	body, err := FetchData(p)
	if err != nil {
		return models.QuoteTrendingResponse{}, err
	}

	var trendingResp models.QuoteTrendingResponse
	if err := json.Unmarshal(body, &trendingResp); err != nil {
		return models.QuoteTrendingResponse{}, err
	}

	return trendingResp, nil
}