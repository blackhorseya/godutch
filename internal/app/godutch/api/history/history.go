package history

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare history api handlers
type IHandler interface {
	// GetByID serve caller to given id to get a record detail
	GetByID(c *gin.Context)

	// List serve caller list all spend records of activity
	List(c *gin.Context)

	// NewRecord serve caller to create a spend record for activity
	NewRecord(c *gin.Context)

	// Delete serve caller to given id to delete a record
	Delete(c *gin.Context)
}

// ProviderSet is provider set for wire
var ProviderSet = wire.NewSet()
