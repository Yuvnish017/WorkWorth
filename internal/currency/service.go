package currency

import (
	"errors"
	"time"
)

type CurrencyService struct {
	currencyRepo *currencyRepository
}

func NewCurrencyService(currencyRepo *currencyRepository) *CurrencyService {
	return &CurrencyService{
		currencyRepo: currencyRepo,
	}
}

func ValidateSupportedCurrency(currency string) bool {
	validCurrencies := map[string]bool{
		"JPY": true,
		"USD": true,
		"EUR": true,
		"INR": true,
	}

	valid, ok := validCurrencies[currency]
	if !ok || !valid {
		return false
	}
	return true
}

func (cs *CurrencyService) ConvertCurrency(request ConversionRequest) (*ConversionResponse, error) {
	if !ValidateSupportedCurrency(request.BaseCurrency) || !ValidateSupportedCurrency(request.TargetCurrency) {
		return nil, errors.New("Either base or target currency not supported")
	}

	exchangeRate, err := cs.currencyRepo.GetRate(request.BaseCurrency, request.TargetCurrency)
	if err != nil {
		return nil, err
	}

	if exchangeRate == nil || time.Since(exchangeRate.UpdatedAt) > 1*time.Hour {
		return nil, errors.New("Old currency rate in DB")
	}

	response := ConversionResponse{
		OriginalValue:  request.Value,
		BaseCurrency:   request.BaseCurrency,
		ConvertedValue: float64(request.Value * exchangeRate.Rate),
		TargetCurrency: request.TargetCurrency,
		Rate:           exchangeRate.Rate,
	}

	return &response, nil
}
