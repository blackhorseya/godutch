package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Activity declare an activity information
type Activity struct {
	ID        int64           `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	Members   []*user.Profile `json:"members"`
	CreatedAt int64           `json:"created_at" db:"created_at"`
}
