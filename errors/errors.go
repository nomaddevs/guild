package errors

import (
	"errors"
)

// Battle.net errors for clients and various packages.
var (
	ErrNotLoggedIn = errors.New("user is not logged in")
)
