package errs

import "errors"

var (
	ErrPostNotFound = errors.New("post by given id is not found")
)
