package currency

import "time"

type ErrorResponse struct {
	Message string
}

type ExchangeRate struct {
	ID             int64
	BaseCurrency   string
	TargetCurrency string
	Rate           float64
	UpdatedAt      time.Time
}

type ConversionRequest struct {
	Value          float64 `json:"value" binding:"required"`
	BaseCurrency   string  `json:"base_currency" binding:"required"`
	TargetCurrency string  `json:"target_currency" binding:"required"`
}

type ConversionResponse struct {
	OriginalValue  float64 `json:"original_value"`
	BaseCurrency   string  `json:"base_currency"`
	ConvertedValue float64 `json:"converted_value"`
	TargetCurrency string  `json:"target_currency"`
	Rate           float64 `json:"rate"`
}
