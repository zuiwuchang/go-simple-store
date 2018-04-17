package dberr

import (
	"errors"
)

// ErrUserNotFound .
var ErrUserNotFound = errors.New("not found user")

// ErrUserCodeNotMatch .
var ErrUserCodeNotMatch = errors.New("active code not match")

// ErrUserDisabled .
var ErrUserDisabled = errors.New("User is disabled")
