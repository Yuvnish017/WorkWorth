package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeProvider struct {
	currencyService *currencyService
}

func NewExchangeProvider(cs *currencyService) *ExchangeProvider {
	return &ExchangeProvider{
		currencyService: cs,
	}
}

func (ep *ExchangeProvider) ConvertCurrency(c *gin.Context) {
	var conversionRequest ConversionRequest

	if err := c.ShouldBind(&conversionRequest); err != nil {
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
