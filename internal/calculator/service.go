package calculator

import (
	"WorkWorth/internal/currency"
	"WorkWorth/internal/users"
	"errors"
)

type calculatorService struct {
	userRepo        users.UserRepository
	currencyService currency.CurrencyService
	pythonClient    PythonClient
}

func NewCalculatorService(ur users.UserRepository, cs currency.CurrencyService, pc PythonClient) *calculatorService {
	return &calculatorService{
		userRepo:        ur,
		currencyService: cs,
		pythonClient:    pc,
	}
}

func (cs *calculatorService) CalculatePurchase(request CalculatePurchaseRequest) (*CalculatePurchaseResponse, error) {
	user, err := cs.userRepo.GetUserByID(request.UserId)
	if err != nil {
		return nil, err
	}

	converted_price, err := cs.currencyService.ConvertCurrency(currency.ConversionRequest{
		Value:          request.Price,
		BaseCurrency:   request.Currency,
		TargetCurrency: user.PreferredCurrency,
	})

	if err != nil {
		return nil, err
	}

	if converted_price.ConvertedValue > user.MonthlySalary {
		return nil, errors.New("Purchase price is greater than montly amount available.")
	}

	hourlyRate := user.MonthlySalary / (user.WorkingDays * user.HoursPerDay)
	requiredWorkHours := converted_price.ConvertedValue / hourlyRate

	response := &CalculatePurchaseResponse{
		Affordable:      true,
		WorkHoursNeeded: requiredWorkHours,
		MonthlyImpact:   user.MonthlySalary - converted_price.ConvertedValue,
		Suggestion:      "Go for it!",
	}

	return response, nil
}
