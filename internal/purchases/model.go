package purchases

import "time"

type Purchase struct {
	ID        int64
	UserId    int64
	Name      string
	Price     float64
	Currency  string
	CreatedAt time.Time
}
