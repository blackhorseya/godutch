package history

import (
	"github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	"github.com/blackhorseya/godutch/internal/pkg/entity/event"
	"github.com/google/wire"
)

// IBiz declare history's service function
type IBiz interface {
	// GetByID serve caller to given id to get a record information
	GetByID(ctx contextx.Contextx, id int64) (info *event.Record, err error)

	// List serve caller to list all records by page and size
	List(ctx contextx.Contextx, page, size int) (infos []*event.Record, err error)

	// NewRecord serve caller to create a new record into spend history
	NewRecord(ctx contextx.Contextx, created *event.Record) (info *event.Record, err error)

	// Delete serve caller to delete a record by id
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
