package strings

func IsASCII(chunk string) bool {

	var result bool = true

	for c := 0; c < len(chunk); c++ {

		var chr = string(chunk[c])

		if chr >= "0" && chr <= "9" {
			// Do Nothing
		} else if chr >= "a" && chr <= "z" {
			// Do Nothing
		} else if chr == "/" || chr == "-" {
			// Do Nothing
		} else if chr == "" || chr == "." {
			result = false
			break
		} else {
			result = false
			break
		}

	}

	return result

}
