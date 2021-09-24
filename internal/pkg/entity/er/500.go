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
	// ErrCreateGoal means create a objective is failure
	ErrCreateGoal = newAPPError(http.StatusInternalServerError, 50020, "create an objective is failure")

	// ErrListGoals means list all objectives is failure
	ErrListGoals = newAPPError(http.StatusInternalServerError, 50021, "list all objectives is failure")

	// ErrCountGoal means count all objectives is failure
	ErrCountGoal = newAPPError(http.StatusInternalServerError, 50023, "count objective is failure")

	// ErrUpdateGoal means update a objective is failure
	ErrUpdateGoal = newAPPError(http.StatusInternalServerError, 50024, "update a objective is failure")

	// ErrGetActivityByID means get activity by id is failure
	ErrGetActivityByID = newAPPError(http.StatusInternalServerError, 50025, "get activity by id is failure")

	// ErrDeleteGoal means delete a objective by id is failure
	ErrDeleteGoal = newAPPError(http.StatusInternalServerError, 50026, "delete a objective by id is failure")
)

var (
	// ErrCreateTask means create a task is failure
	ErrCreateTask = newAPPError(http.StatusInternalServerError, 50030, "create a task is failure")

	// ErrUpdateTask means update a task is failure
	ErrUpdateTask = newAPPError(http.StatusInternalServerError, 50031, "update a task is failure")

	// ErrGetTaskByID means get task by id is failure
	ErrGetTaskByID = newAPPError(http.StatusInternalServerError, 50032, "get task by id is failure")

	// ErrListTasks means list all tasks is failure
	ErrListTasks = newAPPError(http.StatusInternalServerError, 50033, "list all tasks is failure")

	// ErrDeleteTask means delete a task is failure
	ErrDeleteTask = newAPPError(http.StatusInternalServerError, 50034, "delete a task is failure")

	// ErrTaskNotExists means task is not exists
	ErrTaskNotExists = newAPPError(http.StatusNotFound, 50035, "task is not exists")
)

var (
	// ErrGetResultByID means get a key result by id is failure
	ErrGetResultByID = newAPPError(http.StatusInternalServerError, 50040, "get a key result by id is failure")

	// ErrKRNotExists means key result not exists
	ErrKRNotExists = newAPPError(http.StatusNotFound, 50041, "key result not exists")

	// ErrListResults means list all key results is failure
	ErrListResults = newAPPError(http.StatusInternalServerError, 50042, "list all key results is failure")

	// ErrDeleteResult means delete a key result is failure
	ErrDeleteResult = newAPPError(http.StatusInternalServerError, 50043, "delete a key result is failure")

	// ErrUpdateResult means update a key result is failure
	ErrUpdateResult = newAPPError(http.StatusInternalServerError, 50044, "update a key result is failure")

	// ErrCreateResult means create a key result is failure
	ErrCreateResult = newAPPError(http.StatusInternalServerError, 50045, "create a key result is failure")
)

var (
	// ErrDBConnect means db connect is failure
	ErrDBConnect = newAPPError(http.StatusInternalServerError, 50001, "db connect is failure")

	// ErrPing means db ping is failure
	ErrPing = newAPPError(http.StatusInternalServerError, 50002, "db ping is failure")
)
