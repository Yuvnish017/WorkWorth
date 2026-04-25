package calculator

type ErrorResponse struct {
	Message string
}

type CalculatePurchaseRequest struct {
	UserId                   int64
	Price                    float64
	Currency                 string
	HappinessRating          int64
	DailyLifeRelevanceRating int64
	UrgencyRating            int64
}

type CalculatePurchaseResponse struct {
	Affordable               bool
	WorkHourEffortEquivalent float64
	EffortDaysEquivalent     float64
	DisposbaleIncomeLeft     float64
	DisposableIncomeImpact   float64
	MonthsToRecover          float64
	SavingsImpact            float64
	WeightedRatingScore      float64
	Suggestion               string
}

type RecommendationInput struct {
	Price           float64
	WorkHoursNeeded float64
	MonthlyImpact   float64
}

type PythonClient interface {
	GetRecommendation(input RecommendationInput) (string, error)
}
