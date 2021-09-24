package repo

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/google/wire"
)

// IRepo declare activity repository functions
type IRepo interface {
	// GetByID serve caller to get an activity by id
	GetByID(ctx contextx.Contextx, id, userID int64) (info *event.Activity, err error)

	// Create an activity with name and members
	Create(ctx contextx.Contextx, created *event.Activity) (info *event.Activity, err error)

	// AddMembers serve caller to add members into activity
	AddMembers(ctx contextx.Contextx, updated *event.Activity) (info *event.Activity, err error)

	// List serve caller to list all activities
	List(ctx contextx.Contextx, userID int64, limit, offset int) (infos []*event.Activity, err error)

	// Count serve caller to count all activities
	Count(ctx contextx.Contextx, userID int64) (total int, err error)

	// Update serve caller to update an activity information
	Update(ctx contextx.Contextx, updated *event.Activity) (info *event.Activity, err error)

	// Delete serve caller to delete an activity by id
	Delete(ctx contextx.Contextx, id, userID int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
