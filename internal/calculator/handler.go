package calculator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalculatorHandler struct {
	calculatorService *CalculatorService
}

func NewCalculatorHandler(cs *CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{
		calculatorService: cs,
	}
}

func (ch *CalculatorHandler) Suggestor(c *gin.Context) {
	var request CalculatePurchaseRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	response, err := ch.calculatorService.CalculatePurchase(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
