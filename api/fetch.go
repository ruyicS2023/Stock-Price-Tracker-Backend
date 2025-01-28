package api

import (
	"net/http"
	"stock-price-tracker/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/stock", getGlobalQuote)
	r.GET("/stock/daily", getDailySeries)
}

func getGlobalQuote(c *gin.Context) {
	symbol := c.Query("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock symbol is required"})
		return
	}

	quoteResp, err := services.FetchGlobalQuote(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the JSON from the Alpha Vantage API directly or shape it however you like
	c.JSON(http.StatusOK, quoteResp)
}

func getDailySeries(c *gin.Context) {
	symbol := c.Query("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock symbol is required"})
		return
	}

	dailyResp, err := services.FetchTimeSeriesDaily(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the raw daily timeseries or structure it for the frontend
	c.JSON(http.StatusOK, dailyResp)
}
