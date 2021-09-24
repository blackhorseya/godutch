package event

import "github.com/blackhorseya/godutch/internal/pkg/entity/user"

// Activity declare an activity information
type Activity struct {
	ID        int64           `json:"id,omitempty" db:"id"`
	Name      string          `json:"name,omitempty" db:"name"`
	OwnerID   int64           `json:"-" db:"owner_id"`
	Owner     *user.Profile   `json:"owner,omitempty" db:"owner"`
	Members   []*user.Profile `json:"members,omitempty" db:"members"`
	CreatedAt int64           `json:"created_at,omitempty" db:"created_at"`
}
