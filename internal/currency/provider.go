package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeProvider struct {
	currencyService *CurrencyService
}

func NewExchangeProvider(cs *CurrencyService) *ExchangeProvider {
	return &ExchangeProvider{
		currencyService: cs,
	}
}

func (ep *ExchangeProvider) ConvertCurrency(c *gin.Context) {
	var conversionRequest ConversionRequest

	if err := c.ShouldBindJSON(&conversionRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	response, err := ep.currencyService.ConvertCurrency(conversionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
