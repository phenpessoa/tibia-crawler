// Package static provides embeded files.
package static

import "embed"

// TestData embeds all test data.
//go:embed testdata/*
var TestData embed.FS
