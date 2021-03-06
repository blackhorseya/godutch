package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Record declare a cost record information
type Record struct {
	ID        int64          `json:"id" db:"id"`
	Activity  *Activity      `json:"activity,omitempty" db:"activity"`
	Remark    string         `json:"remark,omitempty" db:"remark"`
	Payer     *user.Member   `json:"payer,omitempty" db:"payer"`
	Members   []*user.Member `json:"members,omitempty" db:"members"`
	Total     int            `json:"total" db:"total"`
	CreatedAt int64          `json:"created_at" db:"created_at"`
}
