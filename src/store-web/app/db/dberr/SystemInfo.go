package dberr

import (
	"errors"
)

// ErrSystemInfoEmpty 系統表為空
var ErrSystemInfoEmpty = errors.New("table SystemInfo is empty")

// ErrSystemInfoUnknowRegisterMode .
var ErrSystemInfoUnknowRegisterMode = errors.New("Unknow register mode")

// ErrSystemInfoActiveTitleEmpty .
var ErrSystemInfoActiveTitleEmpty = errors.New("active title can't be empty")

// ErrSystemInfoActiveTextEmpty .
var ErrSystemInfoActiveTextEmpty = errors.New("active text can't be empty")
