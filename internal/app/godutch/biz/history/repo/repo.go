package repo

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/google/wire"
)

// IRepo declare history repository functions
type IRepo interface {
	// GetByID serve caller to get a record by id
	GetByID(ctx contextx.Contextx, id int64) (info *event.Record, err error)

	// List serve caller to list spend history of activity
	List(ctx contextx.Contextx, actID int64, limit, offset int) (infos []*event.Record, err error)

	// Create a new record into activity's spend history
	Create(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error)

	// Delete a record by id
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
