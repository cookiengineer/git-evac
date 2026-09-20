package structs

import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "sync"

type RemoteSettings struct {
	mutex sync.RWMutex
	// "github"
	// map[remote-name]Remote{
	//   Name: "github",
	//   URL:  "git@github.com:{owner}/{repo}.git"
	//   URL:  "https://github.com/{owner}/{repo}.git"
	// }
	Name string `json:"name"`
	URL  string `json:"url"`
	Type string `json:"type"`
}

func NewRemoteSettings(name string) *RemoteSettings {

	var settings RemoteSettings

	settings.Name = name
	settings.URL = ""
	settings.Type = "git"

	return &settings

}

func (settings *RemoteSettings) MarshalJSON() ([]byte, error) {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	type Alias RemoteSettings

	return json.Marshal((*Alias)(settings))

}

func (settings *RemoteSettings) IsValid() bool {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	if utils_strings.IsName(settings.Name) {

		// TODO: Validate URL Schemes

		return true

	}

	return false

}
