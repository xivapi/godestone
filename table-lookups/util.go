package lookups

import "strings"

func listContains(test string, list ...string) bool {
	for _, listItem := range list {
		itemLower := strings.ToLower(listItem)
		if strings.Contains(itemLower, test) {
			return true
		}
	}

	return false
}
