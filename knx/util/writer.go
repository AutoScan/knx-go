// Copyright 2017 Ole Krüger.

package util

import (
	"errors"
)

// BadWriter is a Writer that always fails.
type BadWriter struct{}

// ErrBadWrite is the error that occurs when writing to BadWriter.
var ErrBadWrite = errors.New("Bad write")

// Write implements the Writer Write method.
func (BadWriter) Write([]byte) (int, error) {
	return 0, ErrBadWrite
}
