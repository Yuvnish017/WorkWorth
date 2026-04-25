package calculator

import (
	"errors"

	"WorkWorth/internal/currency"
	"WorkWorth/internal/users"
)

type CalculatorService struct {
	userRepo        users.UserRepository
	currencyService currency.CurrencyService
}

func NewCalculatorService(ur users.UserRepository, cs currency.CurrencyService) *CalculatorService {
	return &CalculatorService{
		userRepo:        ur,
		currencyService: cs,
	}
}

func (cs *CalculatorService) CalculatePurchase(request CalculatePurchaseRequest) (*CalculatePurchaseResponse, error) {
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

	disposableIncome := user.MonthlySalary - user.FixedExpenses - user.EstimateVariableExpenses
	if disposableIncome <= 0 {
		return nil, errors.New("No disposable income available. Check expense inputs.")
	}

	affordable := converted_price.ConvertedValue <= disposableIncome
	hourlyRate := user.MonthlySalary / (float64(user.MonthlyWorkingDays) * user.HoursPerDay)
	workHoursEffortEquivalent := converted_price.ConvertedValue / hourlyRate
	effortDays := workHoursEffortEquivalent / user.HoursPerDay
	impactOnDisposableIncome := (converted_price.ConvertedValue / disposableIncome) * 100
	disposableIncomeLeft := disposableIncome - converted_price.ConvertedValue
	monthsToRecover := converted_price.ConvertedValue / disposableIncome

	savingsImpact := 0.0
	availableAfterSavings := disposableIncome - user.TargetMonthlySaving
	if converted_price.ConvertedValue > availableAfterSavings {
		savingsImpact = converted_price.ConvertedValue - availableAfterSavings
	}

	weightedRatingScore := ((float64(request.HappinessRating) * 0.4) +
		(float64(request.UrgencyRating) * 0.4) +
		(float64(request.DailyLifeRelevanceRating) * 0.2)) / 10.0

	response := &CalculatePurchaseResponse{
		Affordable:               affordable,
		WorkHourEffortEquivalent: workHoursEffortEquivalent,
		EffortDaysEquivalent:     effortDays,
		DisposbaleIncomeLeft:     disposableIncomeLeft,
		DisposableIncomeImpact:   impactOnDisposableIncome,
		MonthsToRecover:          monthsToRecover,
		SavingsImpact:            savingsImpact,
		WeightedRatingScore:      weightedRatingScore,
		Suggestion:               generateResponse(impactOnDisposableIncome, weightedRatingScore, savingsImpact),
	}

	return response, nil
}

func generateResponse(impactOnDisposableIncome, weightedRatingScore, savingsImpact float64) string {
	var suggestion string

	if impactOnDisposableIncome < 20.0 {
		suggestion = "This looks like a low-impact purchase."
	} else if impactOnDisposableIncome < 50.0 {
		suggestion = "This is a moderate purchase. Consider your priorities."
	} else {
		suggestion = "This is a high-impact purchase. You may want to delay."
	}

	if weightedRatingScore > 0.7 {
		suggestion += " However, this seems important to you."
	}

	if savingsImpact > 0.0 {
		suggestion += " This may reduce your planned savings."
	}
	return suggestion
}
