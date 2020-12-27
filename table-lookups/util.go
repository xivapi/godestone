package lookups

import (
	"regexp"
	"strings"
)

func listContains(test string, list ...string) bool {
	for _, listItem := range list {
		itemLower := removeBracketedPhrases(strings.ToLower(listItem))
		if strings.Contains(itemLower, test) {
			return true
		}
	}

	return false
}

var bracketed = regexp.MustCompile("\\[.*\\]")

func removeBracketedPhrases(input string) string {
	return bracketed.ReplaceAllString(input, "")
}
