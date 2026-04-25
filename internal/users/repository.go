package users

import (
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(cur *CreateUserRequest) (*User, error) {
	query := `
		INSERT INTO users (name, email, password, monthly_salary, monthly_working_days, hours_per_day, preferred_currency, fixed_expenses, variable_expenses, target_saving)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, name, email, password, monthly_salary, monthly_working_days, hours_per_day, preferred_currency, fixed_expenses, variable_expenses, target_saving, created_at
	`

	var user User

	err := ur.db.QueryRow(query, cur.Name, cur.Email, cur.Password, cur.MonthlySalary, cur.MonthlyWorkingDays, cur.HoursPerDay, cur.PreferredCurrency, cur.FixedExpenses, cur.EstimateVariableExpenses, cur.TargetMonthlySaving).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.MonthlySalary,
		&user.MonthlyWorkingDays,
		&user.HoursPerDay,
		&user.PreferredCurrency,
		&user.FixedExpenses,
		&user.EstimateVariableExpenses,
		&user.TargetMonthlySaving,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT id, name, email, password, monthly_salary, monthly_working_days, hours_per_day, preferred_currency, fixed_expenses, variable_expenses, target_saving, created_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := ur.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.MonthlySalary,
		&user.MonthlyWorkingDays,
		&user.HoursPerDay,
		&user.PreferredCurrency,
		&user.FixedExpenses,
		&user.EstimateVariableExpenses,
		&user.TargetMonthlySaving,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserByID(id int64) (*User, error) {
	query := `
		SELECT id, name, email, password, monthly_salary, monthly_working_days, hours_per_day, preferred_currency, fixed_expenses, variable_expenses, target_saving, created_at
		FROM users
		WHERE id = $1
	`

	var user User
	err := ur.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.MonthlySalary,
		&user.MonthlyWorkingDays,
		&user.HoursPerDay,
		&user.PreferredCurrency,
		&user.FixedExpenses,
		&user.EstimateVariableExpenses,
		&user.TargetMonthlySaving,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
