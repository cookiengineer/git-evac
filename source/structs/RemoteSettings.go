package structs

import utils_strings "git-evac/utils/strings"

type RemoteSettings struct {
	// "github"
	Name    string             `json:"name"`
	// map[remote-name]Remote{
	//   Name: "github",
	//   URL:  "git@github.com:{owner}/{repo}.git"
	//   URL:  "https://github.com/{owner}/{repo}.git"
	// }
	Remotes map[string]*Remote `json:"remotes"`
	Owners  []string           `json:"owners"`
}

func (settings *RemoteSettings) IsValid() bool {

	if utils_strings.IsName(settings.Name) {

		valid_remotes := true
		valid_owners := true

		for name, remote := range settings.Remotes {

			if !utils_strings.IsName(name) {
				valid_remotes = false
				break
			}

			if !remote.IsValidSchema() {
				valid_remotes = false
				break
			}

		}

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
