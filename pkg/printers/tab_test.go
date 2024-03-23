package printers

import (
	"bytes"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/golangci/golangci-lint/pkg/logutils"
	"github.com/golangci/golangci-lint/pkg/result"
)

func TestTab_Print(t *testing.T) {
	issues := []result.Issue{
		{
			FromLinter: "linter-a",
			Severity:   "warning",
			Text:       "some issue",
			Pos: token.Position{
				Filename: "path/to/filea.go",
				Offset:   2,
				Line:     10,
				Column:   4,
			},
		},
		{
			FromLinter: "linter-b",
			Severity:   "error",
			Text:       "another issue",
			SourceLines: []string{
				"func foo() {",
				"\tfmt.Println(\"bar\")",
				"}",
			},
			Pos: token.Position{
				Filename: "path/to/fileb.go",
				Offset:   5,
				Line:     300,
				Column:   9,
			},
		},
	}

	testCases := []struct {
		desc            string
		printLinterName bool
		useColors       bool
		expected        string
	}{
		{
			desc:            "with linter name",
			printLinterName: true,
			useColors:       false,
			expected: `path/to/filea.go:10:4   linter-a  some issue
path/to/fileb.go:300:9  linter-b  another issue
`,
		},
		{
			desc:            "disable all options",
			printLinterName: false,
			useColors:       false,
			expected: `path/to/filea.go:10:4   some issue
path/to/fileb.go:300:9  another issue
`,
		},
		{
			desc:            "enable all options",
			printLinterName: true,
			useColors:       true,
			expected:        "path/to/filea.go:10:4   linter-a  some issue\npath/to/fileb.go:300:9  linter-b  another issue\n",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			buf := new(bytes.Buffer)

			printer := NewTab(test.printLinterName, test.useColors, logutils.NewStderrLog(logutils.DebugKeyEmpty), buf)

			err := printer.Print(issues)
			require.NoError(t, err)

			assert.Equal(t, test.expected, buf.String())
		})
	}
}
