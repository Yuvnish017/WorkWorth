package currency

import "database/sql"

type currencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) *currencyRepository {
	return &currencyRepository{
		db: db,
	}
}

func (cr *currencyRepository) GetRate(base, target string) (*ExchangeRate, error) {
	query := `
		SELECT id, base_currency, target_currency, rate, updated_at
		FROM currencies
		WHERE base_currency = $1 AND target_currency = $2
	`

	var exchangeRate ExchangeRate

	err := cr.db.QueryRow(query, base, target).Scan(
		&exchangeRate.ID,
		&exchangeRate.BaseCurrency,
		&exchangeRate.TargetCurrency,
		&exchangeRate.Rate,
		&exchangeRate.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func (cr *currencyRepository) SaveRate(base, target string, rate float64) (*ExchangeRate, error) {
	query := `
		INSERT INTO currencies (base_currency, target_currency, rate)
		VALUES ($1, $2, $3)
		RETURNING id, base_currency, target_currency, rate, updates_at
	`

	var exchangeRate ExchangeRate

	err := cr.db.QueryRow(query, base, target, rate).Scan(
		&exchangeRate.ID,
		&exchangeRate.BaseCurrency,
		&exchangeRate.TargetCurrency,
		&exchangeRate.Rate,
		&exchangeRate.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func (cr *currencyRepository) UpdateRate(base, target string, rate float64) (*ExchangeRate, error) {
	query := `
		UPDATE currencies
		SET base_currency = $1, target_currency = $2, rate = $3, updated_at = DEFAULT
		RETURNING id, base_currency, target_currency, rate, updates_at
	`

	var exchangeRate ExchangeRate

	err := cr.db.QueryRow(query, base, target, rate).Scan(
		&exchangeRate.ID,
		&exchangeRate.BaseCurrency,
		&exchangeRate.TargetCurrency,
		&exchangeRate.Rate,
		&exchangeRate.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}
