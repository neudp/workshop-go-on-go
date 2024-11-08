package order

import "errors"

var (
	ErrOrderIsImmutable     = errors.New("order is immutable")
	ErrTransitionNotAllowed = errors.New("transition not allowed")
)
