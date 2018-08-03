package errors

// Error implements the golang error interface.
type Error struct {
	Message string
	Package string
	Type    string
	Method  string
}

// Error returns the error as a string.
func (e *Error) Error() string {
	return e.Message
}

// New returns a new Error.
func New(message string) *Error {
	return &Error{
		Message: message,
	}
}
