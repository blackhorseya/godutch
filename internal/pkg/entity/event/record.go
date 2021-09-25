package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Record declare a cost record information
type Record struct {
	ID         int64          `json:"id" db:"id"`
	Remark     string         `json:"remark,omitempty" db:"remark"`
	PayerID    int64          `json:"–" db:"payer_id"`
	Payer      *user.Member   `json:"payer,omitempty" db:"payer"`
	Members    []*user.Member `json:"members,omitempty" db:"members"`
	TotalSpend int64          `json:"total_spend" db:"total_spend"`
	CreatedAt  int64          `json:"created_at" db:"created_at"`
}
