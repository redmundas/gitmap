package out

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

var _json = `
[
  {
    "count": 1,
    "path": "README.md"
  },
  {
    "count": 2,
    "path": "main.go"
  }
]
`
var _table = `
+-------+-----------+
| Count |   Path    |
+-------+-----------+
| 1     | README.md |
| 2     | main.go   |
+-------+-----------+
`

func TestFormatJson(t *testing.T) {
	output := Json(changes)
	assert.Equal(t, strings.Trim(_json, "\n"), output)
}

func TestFormatTable(t *testing.T) {
	output := Table(changes)
	assert.Equal(t, strings.Trim(_table, "\n"), output)
}
