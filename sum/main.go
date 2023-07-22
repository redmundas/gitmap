package sum

import (
	"path/filepath"
)

type Change struct {
	Count int    `json:"count"`
	Path  string `json:"path"`
}

func HeatMap(lines []string, threshold int, ignoreList []string) []Change {
	heatMap := make(map[string]int)

	for _, line := range lines {
		if line == "" || shouldSkip(line, ignoreList) {
			continue
		}

		if count, ok := heatMap[line]; ok {
			heatMap[line] = count + 1
		} else {
			heatMap[line] = 1
		}
	}

	changes := []Change{}
	for key, val := range heatMap {
		if val < threshold {
			continue
		}
		changes = append(changes, Change{
			Count: val,
			Path:  key,
		})
	}

	return changes
}

func shouldSkip(value string, patterns []string) bool {
	for _, pattern := range patterns {
		match, _ := filepath.Match(pattern, value)
		if match {
			return true
		}
	}
	return false
}
