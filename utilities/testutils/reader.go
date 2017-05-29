package testutils

import (
	"errors"
)

// BadReader is a Reader that always fails.
type BadReader struct{}

// ErrBadRead is the error that occurs when reading from BadReader.
var ErrBadRead = errors.New("Bad read")

// Read implements the Reader Read method.
func (BadReader) Read([]byte) (int, error) {
	return 0, ErrBadRead
}
