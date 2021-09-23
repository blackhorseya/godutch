package activity

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/google/wire"
)

// IBiz declare activity's service function
type IBiz interface {
	// GetByID serve caller to given id to get an activity information
	GetByID(ctx contextx.Contextx, id int64) (info *event.Activity, err error)

	// List serve caller to list all activities by page and size
	List(ctx contextx.Contextx, page, size int) (infos []*event.Activity, total int, err error)

	// NewWithMembers serve caller to create an activity
	NewWithMembers(ctx contextx.Contextx, name string, email []string) (info *event.Activity, err error)

	// ChangeName serve caller to change activity's name
	ChangeName(ctx contextx.Contextx, id int64, name string) (info *event.Activity, err error)

	// Delete serve caller to given id to remove an activity
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
