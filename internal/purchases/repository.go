package purchases

import "database/sql"

type PurchaseRepository interface {
	CreatePurchase(userId int64, name string, price float64, currency string) (*Purchase, error)
	GetPurchasesByUserId(userId int64) ([]*Purchase, error)
}

type purchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *purchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

func (pr *purchaseRepository) CreatePurchase(userId int64, name string, price float64, currency string) (*Purchase, error) {
	query := `
		INSERT INTO purchase (userId, name, price, currency)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, name, price, currency, created_at
	`

	var purchase Purchase

	err := pr.db.QueryRow(query, userId, name, price, currency).Scan(
		&purchase.ID,
		&purchase.UserId,
		&purchase.Name,
		&purchase.Price,
		&purchase.Currency,
		&purchase.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}
