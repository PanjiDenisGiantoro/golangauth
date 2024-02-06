package transaction

import (
	"auth2/user"
	"time"
)

type Transaction struct {
	ID        int    `json:"id"`
	Campaign  int    `json:"campaign"`
	UserID    int    `json:"user_id"`
	Amount    int    `json:"amount"`
	Status    string `json:"status"`
	Code      string `json:"code"`
	User      user.User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
