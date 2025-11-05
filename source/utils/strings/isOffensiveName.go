package strings

import "strings"

func isOffensiveName(name string) bool {

	var result bool

	for _, word := range OffensiveWords {

		if word == name || strings.Contains(name, word) {
			result = true
			break
		}

	}

	return result

}
