package cmd

import (
	"gitmap/sum"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var changes = []sum.Change{
	{Count: 1, Path: "README.md"},
	{Count: 2, Path: "main.go"},
}

var table = `
+-------+-----------+
| Count |   Path    |
+-------+-----------+
| 1     | README.md |
| 2     | main.go   |
+-------+-----------+
`

func TestFormatJson(t *testing.T) {
	output := formatJson(changes)
	assert.Equal(t, "[{\"count\":1,\"path\":\"README.md\"},{\"count\":2,\"path\":\"main.go\"}]", output)
}

func TestFormatTable(t *testing.T) {
	output := formatTable(changes)
	assert.Equal(t, strings.Trim(table, "\n"), output)
}
