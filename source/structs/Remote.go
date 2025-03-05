package structs

import utils_strings "git-evac/utils/strings"
import "strings"

type Remote struct {
	Name string `json:"name"`
	URL  string `json:"url"` // TODO: should be net/url.URL pointer
}

func NewRemote(name string, url string) Remote {

	var remote Remote

	remote.Name = name
	remote.URL  = url

	return remote

}

func (remote *Remote) IsValid() bool {

	name := remote.Name
	url := remote.URL

	if strings.HasPrefix(url, "git@bitbucket.org") ||
		strings.HasPrefix(url, "ssh://git@bitbucket.org") ||
		strings.HasPrefix(url, "https://bitbucket.org/") {

		if utils_strings.IsName(name) {
			return true
		}

	} else if strings.HasPrefix(url, "git@github.com") ||
		strings.HasPrefix(url, "ssh://git@github.com") ||
		strings.HasPrefix(url, "https://github.com/") {

		if utils_strings.IsName(name) {
			return true
		}

	} else if strings.HasPrefix(url, "git@gitlab.com") ||
		strings.HasPrefix(url, "ssh://git@gitlab.com") ||
		strings.HasPrefix(url, "https://gitlab.com/") {

		if utils_strings.IsName(name) {
			return true
		}

	} else if strings.HasPrefix(url, "ssh://git@") ||
		strings.HasPrefix(url, "http://") {

		if utils_strings.IsName(name) {
			return true
		}

	} else if strings.HasPrefix(url, "git://") {

		if utils_strings.IsName(name) {
			return true
		}

	}

	return false

}

func (remote *Remote) IsValidSchema() bool {

	if remote.IsValid() {

		url := remote.URL

		if strings.Contains(url, "{owner}") && strings.Contains(url, "{repository}") {
			return true
		} else if strings.Contains(url, "{") || strings.Contains(url, "}") {
			return false
		}

	}

	return false

}

func (remote *Remote) ToURL(owner string, repository string) string {

	var result string = ""

	if remote.IsValidSchema() {

		tmp := remote.URL
		tmp = strings.ReplaceAll(tmp, "{owner}", owner)
		tmp = strings.ReplaceAll(tmp, "{repository}", repository)

		return tmp

	}

	return result

}
