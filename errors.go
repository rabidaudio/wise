package main

import "errors"

// FriendlyError is a wrapper for errors which
// have a user-friendly explanation of what's going
// on. It is most likely user error, which should
// have some beginner-friendly description of what
// is happening and how to correct it.
//
// Any errors that propigate to the top which is
// not a FriendlyError will be assumed to be an error
// in Wise and will suggest users open an issue.
type FriendlyError struct {
	error
}

func NewFriendlyError(message string) error {
	return FriendlyError{errors.New(message)}
}
