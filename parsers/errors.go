package parsers

import "errors"

var (
	// ErrCtxDone is an error indicating that the context passed to a parser is
	// done.
	ErrCtxDone = errors.New("parsers: ctx done")
)
