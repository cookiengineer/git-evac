package git

import "fmt"

func parseGlobalConfigValue(config *GlobalConfig, section_type string, section_name string, key string, value string) {

	switch section_type {
	case "user":

		switch key {
		case "name":
			config.User.Name = value
		case "email":
			config.User.Email = value
		}

	case "alias":
	case "difftool":
	case "pager":
	case "init":
	}

	fmt.Println("section:", section_type, section_name)
	fmt.Println("\t" + key + ":", "\"" + value + "\"")

}
