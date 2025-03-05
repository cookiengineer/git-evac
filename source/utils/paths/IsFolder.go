package paths

import os_user "os/user"
import "path/filepath"
import "strings"

func IsFolder(path string) bool {

	tmp := strings.TrimSpace(path)

	if strings.HasPrefix(tmp, "~/") {

		user, err := os_user.Current()

		if err == nil {
			tmp = user.HomeDir + "/" + tmp[2:]
		}

	} else if strings.Contains(tmp, "~") {
		tmp = ""
	}

	if strings.HasPrefix(tmp, "/") {

		resolved, err := filepath.Abs(tmp)

		if err == nil && resolved == tmp {
			return true
		}

	}

	return false

}
