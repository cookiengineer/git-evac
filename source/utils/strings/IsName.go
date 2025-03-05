package strings

func IsName(name string) bool {

	var result bool = true

	if len(name) > 16 {
		result = false
	}

	if result == true {

		for n := 0; n < len(name); n++ {

			chr := string(name[n])
			valid := false

			if chr >= "a" && chr <= "z" {
				valid = true
			} else if chr >= "0" && chr <= "9" {
				valid = true
			} else if chr == "-" {
				valid = true
			} else if chr == "_" {
				valid = true
			}

			if valid == false {
				result = false
				break
			}

		}

	}

	if result == true {

		var normalized string

		for n := 0; n < len(name); n++ {

			chr := string(name[n])

			if chr >= "a" && chr <= "z" {
				normalized = normalized + chr
			} else if chr >= "0" && chr <= "9" {
				normalized = normalized + chr
			}

		}

		if isOffensiveName(normalized) == true {
			result = false
		}

	}

	return result

}
