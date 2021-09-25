package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Activity declare an activity information
type Activity struct {
	ID        int64          `json:"id,omitempty" db:"id"`
	Name      string         `json:"name,omitempty" db:"name"`
	Members   []*user.Member `json:"members,omitempty" db:"members"`
	CreatedAt int64          `json:"created_at,omitempty" db:"created_at"`
}
