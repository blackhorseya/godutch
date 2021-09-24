package activity

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare activity api handlers
type IHandler interface {
	// GetByID serve caller to given id to get an activity information
	GetByID(c *gin.Context)

	// List serve caller to list all activities
	List(c *gin.Context)

	// NewWithMembers serve caller to create a new activity with members email
	NewWithMembers(c *gin.Context)

	// ChangeName serve caller to change activity's name by id
	ChangeName(c *gin.Context)

	// Delete serve caller to delete an activity by id
	Delete(c *gin.Context)
}

// ProviderSet is provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
