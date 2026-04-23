package users

import (
	"time"
)

type User struct {
	ID                       int64
	Name                     string
	Email                    string
	Password                 string
	MonthlySalary            float64
	MonthlyWorkingDays       int64
	HoursPerDay              float64
	PreferredCurrency        string
	FixedExpenses            float64
	EstimateVariableExpenses float64
	CreatedAt                time.Time
}

type CreateUserRequest struct {
	Name                     string
	Email                    string
	Password                 string
	MonthlySalary            float64
	MonthlyWorkingDays       int64
	HoursPerDay              float64
	PreferredCurrency        string
	FixedExpenses            float64
	EstimateVariableExpenses float64
}

type UserRepository interface {
	Create(*CreateUserRequest) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int64) (*User, error)
}
