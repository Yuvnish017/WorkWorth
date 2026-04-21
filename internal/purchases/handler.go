package purchases

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	purchaseService *PurchaseService
}

func NewPurchaseHandler(ps *PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseService: ps,
	}
}

func (ph *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var request CreatePurchaseRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userId, err := strconv.ParseInt(c.GetString("x-user-id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Issue in extracting user id"})
		return
	}
	fmt.Println("Sending to service...")
	response, err := ph.purchaseService.CreatePurchase(userId, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (ph *PurchaseHandler) FetchPurchasesByUserId(c *gin.Context) {
	userId, err := strconv.ParseInt(c.GetString("x-user-id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Issue in extracting user id"})
		return
	}
	response, err := ph.purchaseService.GetPurchaseByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
