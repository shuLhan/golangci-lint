package logutils

import (
	"os"
)

var (
	StdOut = os.Stdout // https://github.com/golangci/golangci-lint/issues/14
	StdErr = os.Stderr
)
