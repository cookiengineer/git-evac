package structs

import utils_strings "git-evac/utils/strings"

type OrganizationSettings struct {
	Name       string                      `json:"name"`
	Identities map[string]IdentitySettings `json:"identities"`
	Remotes    map[string]RemoteSettings   `json:"remotes"`
}

func (settings *OrganizationSettings) IsValid() bool {

	valid_identities := true
	valid_remotes := true

	for name, identity := range settings.Identities {

		if utils_strings.IsName(name) && identity.IsValid() == false {
			valid_identities = false
			break
		}

	}

	for name, remote := range settings.Remotes {

		if utils_strings.IsName(name) && remote.IsValid() == false {
			valid_remotes = false
			break
		}

	}

	return valid_identities && valid_remotes

}
