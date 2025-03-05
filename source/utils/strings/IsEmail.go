package strings

import "strings"

func IsEmail(value string) bool {

	if strings.Contains(value, "@") {

		tmp := strings.Split(value, "@")

		if len(tmp) == 2 {

			if IsName(tmp[0]) && IsDomain(tmp[1]) {
				return true
			}

		}

	}

	return false

}
