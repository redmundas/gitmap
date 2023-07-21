package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatMap(t *testing.T) {
	input := []string{"main.go", "README.md", "main.go"}
	output := HeatMap(input, 1, []string{})
	assert.Equal(t, []Change{
		{Count: 2, Path: "main.go"},
		{Count: 1, Path: "README.md"},
	}, output)
}

func TestShouldSkip(t *testing.T) {
	assert.Equal(t, true, shouldSkip("README.md", []string{".md"}))
	assert.Equal(t, false, shouldSkip("main.go", []string{".md"}))
}
