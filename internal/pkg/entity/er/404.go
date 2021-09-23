package er

import "net/http"

var (
	// ErrUserNotExists means user is not exists
	ErrUserNotExists = newAPPError(http.StatusNotFound, 40400, "user is not exists")

	// ErrGoalNotExists means goal not exists
	ErrGoalNotExists = newAPPError(http.StatusNotFound, 40401, "goal not exists")
)
