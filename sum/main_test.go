package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatMap(t *testing.T) {
	input := []string{"main.go", "README.md", "main.go", "go.mod", "go.sum"}
	output := HeatMap(input, 1, false, []string{"go.mod", "go.sum"})
	assert.EqualValues(t, []Change{
		{Count: 1, Path: "README.md"},
		{Count: 2, Path: "main.go"},
	}, output)
}

func TestShouldSkip(t *testing.T) {
	assert.Equal(t, true, shouldSkip("docs/README.md", []string{"**/*.md"}))
	assert.Equal(t, false, shouldSkip("docs/README.md", []string{"*.md"}))
	assert.Equal(t, true, shouldSkip("README.md", []string{"*.md"}))
	assert.Equal(t, false, shouldSkip("README.md", []string{"**/*.md"}))
	assert.Equal(t, false, shouldSkip("main.go", []string{"*.md"}))
	assert.Equal(t, false, shouldSkip("go.mod", []string{"go.sum"}))
	assert.Equal(t, true, shouldSkip("go.sum", []string{"go.sum"}))
}
