package purchases

import (
	"errors"
	"strings"
)

type PurchaseService struct {
	purchaseRepo PurchaseRepository
}

func NewPurchaseService(repo PurchaseRepository) *PurchaseService {
	return &PurchaseService{
		purchaseRepo: repo,
	}
}

func (ps *PurchaseService) CreatePurchase(userId int64, request CreatePurchaseRequest) (*Purchase, error) {
	validCurrencies := map[string]bool{
		"JPY": true,
		"USD": true,
		"EUR": true,
	}

	if strings.TrimSpace(request.Name) == "" {
		return nil, errors.New("A name for purchase is required")
	}

	if request.Price <= 0 {
		return nil, errors.New("Price should not be less than or equal to 0")
	}

	if ok := validCurrencies[request.Currency]; !ok {
		return nil, errors.New("Currency not supported")
	}

	purchase, err := ps.purchaseRepo.Create(userId, request.Name, request.Price, request.Currency)
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
