package types

import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "strings"
import "sync"

type Remote struct {
	mutex sync.RWMutex
	Name  string `json:"name"`
	URL   string `json:"url"` // TODO: should be net/url.URL pointer
}

func NewRemote(name string, url string) *Remote {

	var remote Remote

	remote.Name = name
	remote.URL  = url

	return &remote

}

func (remote *Remote) MarshalJSON() ([]byte, error) {

	remote.mutex.RLock()
	defer remote.mutex.RUnlock()

	type Alias Remote

	return json.Marshal((*Alias)(remote))

}

func (remote *Remote) GetName() string {

	var result string

	remote.mutex.RLock()
	result = remote.Name
	remote.mutex.RUnlock()

	return result

}

func (remote *Remote) GetURL() string {

	var result string

	remote.mutex.RLock()
	result = remote.URL
	remote.mutex.RUnlock()

	return result

}

func (remote *Remote) SetURL(value string) {

	remote.mutex.Lock()
	remote.URL = value
	remote.mutex.Unlock()

}

func (remote *Remote) SetName(value string) {

	remote.mutex.Lock()
	remote.Name = value
	remote.mutex.Unlock()

}

func (remote *Remote) IsValid() bool {

	var result bool

	remote.mutex.RLock()

	result = remote.isValidLocked()

	remote.mutex.RUnlock()

	return result

}

func (remote *Remote) isValidLocked() bool {

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

	var result bool

	remote.mutex.RLock()

	result = remote.isValidSchemaLocked()

	remote.mutex.RUnlock()

	return result

}

func (remote *Remote) isValidSchemaLocked() bool {

	if remote.isValidLocked() {

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

	remote.mutex.RLock()

	if remote.isValidSchemaLocked() {

		tmp := remote.URL
		tmp = strings.ReplaceAll(tmp, "{owner}", owner)
		tmp = strings.ReplaceAll(tmp, "{repository}", repository)

		result = tmp

	}

	remote.mutex.RUnlock()

	return result

}
