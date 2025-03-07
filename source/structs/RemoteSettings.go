package structs

import utils_strings "git-evac/utils/strings"

type RemoteSettings struct {
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

func (settings *RemoteSettings) IsValid() bool {

	if utils_strings.IsName(settings.Name) {

		// TODO: Validate URL Schemes

		return true

	}
	
	return false

}
