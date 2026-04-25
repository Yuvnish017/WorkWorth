package purchases

import "time"

type Purchase struct {
	ID        int64
	UserId    int64
	Name      string
	Price     float64
	Currency  string
	CreatedAt time.Time
}

type CreatePurchaseRequest struct {
	Name              string  `form:"name" binding:"required"`
	Price             float64 `form:"price" binding:"required"`
	Currency          string  `form:"currency" binding:"required"`
	PreferredCurrency string  `form:"pref_currency" binding:"required"`
}

type ErrorResponse struct {
	Message string
}
