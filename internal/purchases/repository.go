package purchases

import "database/sql"

type PurchaseRepository interface {
	Create(userId int64, name string, price float64, currency string) (*Purchase, error)
	FetchByUserId(userId int64) ([]*Purchase, error)
}

type purchaseRepository struct {
	db *sql.DB
}

func NewPurchaseRepository(db *sql.DB) *purchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

func (pr *purchaseRepository) Create(userId int64, name string, price float64, currency string) (*Purchase, error) {
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

func (pr *purchaseRepository) FetchByUserId(userId int64) ([]*Purchase, error) {
	query := `
		SELECT id, user_id, name, price, currency, created_at
		FROM purchases
		WHERE user_id = $1
	`

	rows, err := pr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var purchases []*Purchase

	for rows.Next() {
		var p Purchase

		err = rows.Scan(
			&p.ID,
			&p.UserId,
			&p.Name,
			&p.Price,
			&p.Currency,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		purchases = append(purchases, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return purchases, nil
}
