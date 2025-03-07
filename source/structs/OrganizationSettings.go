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

func (settings *OrganizationSettings) GetIdentity(name string) *IdentitySettings {

	var result *IdentitySettings = nil

	if name != "" {

		tmp, ok := settings.Identities[name]

		if ok == true {
			result = &tmp
		}

	}

	return result

}

func (settings *OrganizationSettings) GetRemote(name string) *RemoteSettings {

	var result *RemoteSettings = nil

	if name != "" {

		tmp, ok := settings.Remotes[name]

		if ok == true {
			result = &tmp
		}

	}

	return result

}

func (settings *OrganizationSettings) RemoveIdentity(name string) bool {

	var result bool = false

	if name != "" {

		_, ok := settings.Identities[name]

		if ok == true {
			delete(settings.Identities, name)
			result = true
		}

	}

	return result

}

func (settings *OrganizationSettings) RemoveRemote(name string) bool {

	var result bool = false

	if name != "" {

		_, ok := settings.Remotes[name]

		if ok == true {
			delete(settings.Remotes, name)
			result = true
		}

	}

	return result

}

func (settings *OrganizationSettings) SetIdentity(value IdentitySettings) bool {

	var result bool = false

	if value.Name != "" {
		settings.Identities[value.Name] = value
		result = true
	}

	return result

}

func (settings *OrganizationSettings) SetRemote(value RemoteSettings) bool {

	var result bool = false

	if value.Name != "" {
		settings.Remotes[value.Name] = value
		result = true
	}

	return result

}

