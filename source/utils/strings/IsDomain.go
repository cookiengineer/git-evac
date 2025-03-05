package strings

func IsDomain(value string) bool {

	var result bool = true

	for v := 0; v < len(value); v++ {

		chr := string(value[v])
		valid := false

		if chr >= "a" && chr <= "z" {
			valid = true
		} else if chr >= "0" && chr <= "9" {
			valid = true
		} else if chr == "-" {
			valid = true
		} else if chr == "_" {
			valid = true
		} else if chr == "." {
			valid = true
		}

		if valid == false {
			result = false
			break
		}

	}

	return result

}
