package transaction

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Campaign  int       `json:"campaign"`
	UserID    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
