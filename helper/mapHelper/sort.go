package mapHelper

import (
	"sort"
)

func SortedKeys(m *map[string]string) []string {
	sortedKeys := make([]string, len(*m))
	i := 0
	for key := range *m {
		sortedKeys[i] = key
		i++
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func SortedUrl(m *map[string]string) string {
	if len(*m) == 0 {
		return ""
	}
	sk := SortedKeys(m)
	var sortedData string
	for _, k := range sk {
		sortedData += "&" + k + "=" + (*m)[k]
	}
	return sortedData[1:]
}
