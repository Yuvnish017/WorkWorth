package purchases

import (
	"WorkWorth/internal/currency"
	"errors"
	"strings"
)

type PurchaseService struct {
	purchaseRepo    PurchaseRepository
	currencyService *currency.CurrencyService
}

func NewPurchaseService(repo PurchaseRepository, cs *currency.CurrencyService) *PurchaseService {
	return &PurchaseService{
		purchaseRepo:    repo,
		currencyService: cs,
	}
}

func (ps *PurchaseService) CreatePurchase(userId int64, request CreatePurchaseRequest) (*Purchase, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, errors.New("A name for purchase is required")
	}

	if request.Price <= 0 {
		return nil, errors.New("Price should not be less than or equal to 0")
	}

	conversionResponse, err := ps.currencyService.ConvertCurrency(currency.ConversionRequest{
		Value:          request.Price,
		BaseCurrency:   request.Currency,
		TargetCurrency: request.PreferredCurrency,
	})

	if err != nil {
		return nil, errors.New("Error in currency validation and conversion")
	}

	purchase, err := ps.purchaseRepo.Create(userId, request.Name, conversionResponse.ConvertedValue, conversionResponse.TargetCurrency)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (ps *PurchaseService) GetPurchaseByUserId(userId int64) ([]*Purchase, error) {
	purchases, err := ps.purchaseRepo.FetchByUserId(userId)
	if err != nil {
		return nil, err
	}

	return purchases, err
}
