package dberr

import (
	"errors"
)

// ErrSystemInfoEmpty 系統表為空
var ErrSystemInfoEmpty = errors.New("table SystemInfo is empty")
