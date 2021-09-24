package er

import "net/http"

var (
	// ErrGetUserByID means get user by id is failure
	ErrGetUserByID = newAPPError(http.StatusInternalServerError, 50000, "get user by id is failure")

	// ErrGetUserByToken means get user by token is failure
	ErrGetUserByToken = newAPPError(http.StatusInternalServerError, 50001, "get user by token is failure")

	// ErrGetUserByEmail means get user by email is failure
	ErrGetUserByEmail = newAPPError(http.StatusInternalServerError, 50001, "get user by email is failure")

	// ErrEncryptPassword means encrypt password is failure
	ErrEncryptPassword = newAPPError(http.StatusInternalServerError, 50002, "encrypt password is failure")

	// ErrSignup means signup is failure
	ErrSignup = newAPPError(http.StatusInternalServerError, 50003, "signup is failure")

	// ErrNewToken means new a jwt is failure
	ErrNewToken = newAPPError(http.StatusInternalServerError, 50004, "new a jwt is failure")

	// ErrUpdateToken means update token by user id is failure
	ErrUpdateToken = newAPPError(http.StatusInternalServerError, 50005, "update token by user id is failure")

	// ErrLogin means login is failure
	ErrLogin = newAPPError(http.StatusInternalServerError, 50005, "login is failure")

	// ErrValidateToken means couldn't parse claims
	ErrValidateToken = newAPPError(http.StatusInternalServerError, 50006, "Couldn't parse claims")
)

var (
	// ErrReadiness means readiness is failure
	ErrReadiness = newAPPError(http.StatusInternalServerError, 50010, "readiness is failure")

	// ErrLiveness means liveness is failure
	ErrLiveness = newAPPError(http.StatusInternalServerError, 50011, "liveness is failure")
)

var (
	// ErrCreateActivity means create an activity is failure
	ErrCreateActivity = newAPPError(http.StatusInternalServerError, 50020, "create an activity is failure")

	// ErrListActivities means list all activities is failure
	ErrListActivities = newAPPError(http.StatusInternalServerError, 50021, "list all activities is failure")

	// ErrCountActivity means count all activities is failure
	ErrCountActivity = newAPPError(http.StatusInternalServerError, 50023, "count activities is failure")

	// ErrUpdateActivity means update an activity is failure
	ErrUpdateActivity = newAPPError(http.StatusInternalServerError, 50024, "update an activity is failure")

	// ErrGetActivityByID means get activity by id is failure
	ErrGetActivityByID = newAPPError(http.StatusInternalServerError, 50025, "get an activity by id is failure")

	// ErrDeleteActivity means delete an activity by id is failure
	ErrDeleteActivity = newAPPError(http.StatusInternalServerError, 50026, "delete an activity by id is failure")
)

var (
	// ErrDBConnect means db connect is failure
	ErrDBConnect = newAPPError(http.StatusInternalServerError, 50001, "db connect is failure")

	// ErrPing means db ping is failure
	ErrPing = newAPPError(http.StatusInternalServerError, 50002, "db ping is failure")
)
