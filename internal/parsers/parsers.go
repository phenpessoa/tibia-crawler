// Package parsers provides interfaces and implementations for parsing content
// from tibia.com.
package parsers

import (
	"context"
	"net/http"
)

// Parser defines an interface for parsing content from a specific tibia.com
// page.
//
// A parser is responsible for extracting relevant information from the HTML
// content of a tibia.com page, and transforming it into a structured format
// that can be utilized by other components of the application.
//
// Each parser implementation is specific to a particular type of content on
// tibia.com, such as characters, worlds, etc.
type Parser[A, P any] interface {
	// Parse parses the HTML content from a tibia.com page and returns the
	// parsed data.
	//
	// The Parse method should extract the relevant information from the HTML
	// content, perform any necessary transformations, and return the parsed
	// data as a structured representation, along with any potential error
	// encountered during parsing.
	//
	// If parsing is successful, the parsed data should be returned along
	// with a nil error.
	// If parsing fails, an error should be returned, describing the reason for
	// the failure.
	//
	// Type A should be a type that passes any relevant information to the
	// parser. An example would be a character name.
	Parse(ctx context.Context, args A, opts Options) (parsed P, err error)

	// URL returns the tibia.com endpoint being parsed.
	//
	// The URL method should return a string representing the specific tibia.com
	// endpoint from which the parser is extracting content. This information
	// helps the caller identify the purpose and context of the parser.
	URL() string
}

// Options provides configurable options for the parser.
//
// The Options struct allows customization of the parser's behavior and
// parameters. It can be used to pass additional configuration settings that
// affect the parsing process.
type Options struct {
	// HTTPClient specifies an optional HTTP client to be used
	// for making requests to tibia.com.
	//
	// If an HTTP client is provided, the parser will utilize it for making HTTP
	// requests to fetch the HTML content of the tibia.com page being parsed.
	//
	// If no HTTP client is provided, the parser falls back to
	// the default HTTP client.
	HTTPClient *http.Client
}
