package parsers

import "errors"

var (
	// ErrCtxDone is an error indicating that the context passed to a parser is
	// done.
	ErrCtxDone = errors.New("parsers: ctx done")

	// ErrRateLimited will be sent by parsers in case the requets to tibia.com
	// failed due to ratelimit.
	ErrRateLimited = errors.New("parsers: ratelimited")

	// ErrMaintenance will be sent by parsers in case tibia.com is under
	// maintenance.
	ErrMaintenance = errors.New("parsers: tibia.com is under maintenance")

	// ErrUnknownStatusCode will be sent by parsers in case tibia.com responded
	// with an unknown status code.
	ErrUnknownStatusCode = errors.New("parsers: unknown status code")
)
