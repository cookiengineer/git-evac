package components

import "strings"

func toRemoteLabel(name string, url string) string {

	// TODO: This method needs to use the App Settings later

	var result string

	if name == "bitbucket" ||
		strings.HasPrefix(url, "git@bitbucket.org") ||
		strings.HasPrefix(url, "ssh://git@bitbucket.org") ||
		strings.HasPrefix(url, "https://bitbucket.org/") {
		result = "bitbucket"
	} else if name == "github" ||
		strings.HasPrefix(url, "git@github.com") ||
		strings.HasPrefix(url, "ssh://git@github.com") ||
		strings.HasPrefix(url, "https://github.com/") {
		result = "github"
	} else if name == "gitlab" ||
		strings.HasPrefix(url, "git@gitlab.com") ||
		strings.HasPrefix(url, "ssh://git@gitlab.com") ||
		strings.HasPrefix(url, "https://gitlab.com/") {
		result = "gitlab"
	} else if name == "gitea" {
		result = "gitea"
	} else if name == "gogs" ||
		strings.HasPrefix(url, "ssh://git@homeserver") ||
		strings.HasPrefix(url, "http://homeserver/") {
		result = "gogs"
	} else if strings.HasPrefix(url, "git://") {
		result = "git"
	}

	return result

}
