package fnp

import (
	"errors"
)

type FilenamePattern interface {
	Parse(filename string) (filenameInfo interface{}, err error)
	Format(filenameInfo interface{}) (filename string, err error)
}

var (
	ErrNilFilenamePattern      error = errors.New("FilenamePattern is nil")
	ErrNilFilenameInfo         error = errors.New("FilenameInfo is nil")
	ErrEmptyFilename           error = errors.New("filename is empty")
	ErrFilenameNotMatchPattern error = errors.New(
		"filename does NOT match the pattern")
	ErrFilenameInfoNotMatchPattern error = errors.New(
		"FilenameInfo does NOT match the pattern")
)
