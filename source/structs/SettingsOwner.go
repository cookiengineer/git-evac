package structs

import utils_strings "git-evac/utils/strings"

type SettingsOwner struct {
	Name       string                      `json:"name"`
	Identities map[string]IdentitySettings `json:"identities"`
	Remotes    map[string]RemoteSettings   `json:"remotes"`
	Services   map[string]ServiceSettings  `json:"services"`
}

func (settings *SettingsOwner) IsValid() bool {

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

func (settings *SettingsOwner) GetIdentity(name string) *IdentitySettings {

	var result *IdentitySettings = nil

	if name != "" {

		tmp, ok := settings.Identities[name]

		if ok == true {
			result = &tmp
		}

	}

	return result

}

func (settings *SettingsOwner) GetRemote(name string) *RemoteSettings {

	var result *RemoteSettings

	if name != "" {

		tmp, ok := settings.Remotes[name]

		if ok == true {
			result = &tmp
		}

	}

	return result

}

func (settings *SettingsOwner) GetService(name string) *ServiceSettings {

	var result *ServiceSettings

	if name != "" {

		tmp, ok := settings.Services[name]

		if ok == true {
			result = &tmp
		}

	}

	return result

}

func (settings *SettingsOwner) RemoveIdentity(name string) bool {

	var result bool

	if name != "" {

		_, ok := settings.Identities[name]

		if ok == true {
			delete(settings.Identities, name)
			result = true
		}

	}

	return result

}

func (settings *SettingsOwner) RemoveRemote(name string) bool {

	var result bool

	if name != "" {

		_, ok := settings.Remotes[name]

		if ok == true {
			delete(settings.Remotes, name)
			result = true
		}

	}

	return result

}

func (settings *SettingsOwner) RemoveService(name string) bool {

	var result bool

	if name != "" {

		_, ok := settings.Services[name]

		if ok == true {
			delete(settings.Services, name)
			result = true
		}

	}

	return result

}

func (settings *SettingsOwner) SetIdentity(value IdentitySettings) bool {

	var result bool

	if value.Name != "" {
		settings.Identities[value.Name] = value
		result = true
	}

	return result

}

func (settings *SettingsOwner) SetRemote(value RemoteSettings) bool {

	var result bool

	if value.Name != "" {
		settings.Remotes[value.Name] = value
		result = true
	}

	return result

}

func (settings *SettingsOwner) SetService(value ServiceSettings) bool {

	var result bool

	if value.Name != "" {
		settings.Services[value.Name] = value
		result = true
	}

	return result

}

