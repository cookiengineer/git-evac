package structs

import utils_strings "git-evac/utils/strings"

type RemoteSettings struct {
	// "github"
	// map[remote-name]Remote{
	//   Name: "github",
	//   URL:  "git@github.com:{owner}/{repo}.git"
	//   URL:  "https://github.com/{owner}/{repo}.git"
	// }
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Type   string   `json:"type"`
	Owners []string `json:"owners"`
}

func (settings *RemoteSettings) IsValid() bool {

	if utils_strings.IsName(settings.Name) {

		valid_remotes := true
		valid_owners := true

		for o := 0; o < len(settings.Owners); o++ {

			if !utils_strings.IsName(settings.Owners[o]) {
				valid_owners = false
				break
			}

		}

		return valid_remotes && valid_owners

	}
	
	return false

}
