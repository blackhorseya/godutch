package er

import "net/http"

var (
	// ErrUserNotExists means user is not exists
	ErrUserNotExists = newAPPError(http.StatusNotFound, 40400, "user is not exists")

	// ErrActivityNotExists means activity not exists
	ErrActivityNotExists = newAPPError(http.StatusNotFound, 40401, "activity not exists")
)
