package sum

import (
	"strings"
)

type Change struct {
	Count int
	Path  string
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

func shouldSkip(value string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(value, suffix) {
			return true
		}
	}
	return false
}
