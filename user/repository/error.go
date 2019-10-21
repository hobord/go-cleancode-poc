package repository

// Error is the error type of repository
type Error struct {
	Query string
	Err   error
}

func (e *Error) Unwrap() error {
	return e.Err
}
