package calculator

type CalculatePurchaseRequest struct {
	UserId   int64
	Price    float64
	Currency string
}

type CalculatePurchaseResponse struct {
	Affordable      bool
	WorkHoursNeeded float64
	MonthlyImpact   float64
	Suggestion      string
}

type RecommendationInput struct {
	Price           float64
	WorkHoursNeeded float64
	MonthlyImpact   float64
}

type PythonClient interface {
	GetRecommendation(input RecommendationInput) (string, error)
}
