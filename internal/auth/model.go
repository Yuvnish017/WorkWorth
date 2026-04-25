package auth

type ErrorResponse struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Name                     string  `json:"name" binding:"required"`
	Email                    string  `json:"email" binding:"required,email"`
	Password                 string  `json:"password" binding:"required"`
	MonthlySalary            float64 `json:"monthly_salary" binding:"required"`
	MonthlyWorkingDays       int64   `json:"monthly_working_days" binding:"required"`
	HoursPerDay              float64 `json:"hours_per_day" binding:"required"`
	PreferredCurrency        string  `json:"preferred_currency" binding:"required"`
	FixedExpenses            float64 `json:"fixed_expenses" binding:"required"`
	EstimateVariableExpenses float64 `json:"variable_expenses" binding:"required"`
	TargetMonthlySaving      float64 `json:"target_saving" binding:"required"`
	ConfirmPassword          string  `json:"confirmPassword" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
