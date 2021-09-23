package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Activity declare an activity information
type Activity struct {
	ID        int64           `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	OwnerID   int64           `json:"-" db:"owner_id"`
	Owner     *user.Profile   `json:"owner" db:"owner"`
	Members   []*user.Profile `json:"members" db:"members"`
	CreatedAt int64           `json:"created_at" db:"created_at"`
}
