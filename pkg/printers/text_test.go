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

func TestText_Print(t *testing.T) {
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
		printIssuedLine bool
		printLinterName bool
		expected        string
	}{
		{
			desc:            "printIssuedLine and printLinterName",
			printIssuedLine: true,
			printLinterName: true,
			expected: `path/to/filea.go:10:4: some issue (linter-a)
path/to/fileb.go:300:9: another issue (linter-b)
func foo() {
	fmt.Println("bar")
}
`,
		},
		{
			desc:            "printLinterName only",
			printIssuedLine: false,
			printLinterName: true,
			expected: `path/to/filea.go:10:4: some issue (linter-a)
path/to/fileb.go:300:9: another issue (linter-b)
`,
		},
		{
			desc:            "printIssuedLine only",
			printIssuedLine: true,
			printLinterName: false,
			expected: `path/to/filea.go:10:4: some issue
path/to/fileb.go:300:9: another issue
func foo() {
	fmt.Println("bar")
}
`,
		},
		{
			desc:            "enable all options",
			printIssuedLine: true,
			printLinterName: true,
			expected:        "path/to/filea.go:10:4: some issue (linter-a)\npath/to/fileb.go:300:9: another issue (linter-b)\nfunc foo() {\n\tfmt.Println(\"bar\")\n}\n",
		},
		{
			desc:            "disable all options",
			printIssuedLine: false,
			printLinterName: false,
			expected: `path/to/filea.go:10:4: some issue
path/to/fileb.go:300:9: another issue
`,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			buf := new(bytes.Buffer)

			printer := NewText(test.printIssuedLine, false, test.printLinterName, logutils.NewStderrLog(logutils.DebugKeyEmpty), buf)

			err := printer.Print(issues)
			require.NoError(t, err)

			assert.Equal(t, test.expected, buf.String())
		})
	}
}
