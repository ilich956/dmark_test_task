package errors

import "errors"

var (
	ErrEmptyTitle = errors.New("title is required")
)
