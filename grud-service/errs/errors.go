package errs

import "errors"

var (
	ErrPostNotFound = errors.New("pst by given id is not found")
)
