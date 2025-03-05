package strings

import "strings"

func isOffensiveName(name string) bool {

	var result bool = false

	for o := 0; o < len(OffensiveWords); o++ {

		word := OffensiveWords[o]

		if word == name || strings.Contains(name, word) {
			result = true
			break
		}

	}

	return result

}
