package auth

type ErrorResponse struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Name                     string  `form:"name" binding:"required"`
	Email                    string  `form:"email" binding:"required,email"`
	Password                 string  `form:"password" binding:"required"`
	MonthlySalary            float64 `form:"monthly_salary" binding:"required"`
	MonthlyWorkingDays       int64   `form:"monthly_working_days" binding:"required"`
	HoursPerDay              float64 `form:"hours_per_day" binding:"required"`
	PreferredCurrency        string  `form:"preferred_currency" binding:"required"`
	FixedExpenses            float64 `form:"fixed_expenses" binding:"required"`
	EstimateVariableExpenses float64 `form:"variable_expenses" binding:"required"`
	ConfirmPassword          string  `form:"confirmPassword" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
