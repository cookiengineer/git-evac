package structs

import "git-evac/types"
import utils_strings "git-evac/utils/strings"

type SettingsOwner struct {
	Name       string                     `json:"name"`
	Identities map[string]*types.Identity `json:"identities"`
	Remotes    map[string]*types.Remote   `json:"remotes"`
	Services   map[string]*types.Service  `json:"services"`
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

func (settings *SettingsOwner) GetIdentity(name string) *types.Identity {

	var result *types.Identity = nil

	if name != "" {

		identity, ok := settings.Identities[name]

		if ok == true {
			result = identity
		}

	}

	return result

}

func (settings *SettingsOwner) GetRemote(name string) *types.Remote {

	var result *types.Remote = nil

	if name != "" {

		remote, ok := settings.Remotes[name]

		if ok == true {
			result = remote
		}

	}

	return result

}

func (settings *SettingsOwner) GetService(name string) *types.Service {

	var result *types.Service = nil

	if name != "" {

		service, ok := settings.Services[name]

		if ok == true {
			result = service
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

func (settings *SettingsOwner) SetIdentity(value types.Identity) bool {

	var result bool

	if value.Name != "" {
		settings.Identities[value.Name] = &value
		result = true
	}

	return result

}

func (settings *SettingsOwner) SetRemote(value types.Remote) bool {

	var result bool

	if value.Name != "" {
		settings.Remotes[value.Name] = &value
		result = true
	}

	return result

}

func (settings *SettingsOwner) SetService(value types.Service) bool {

	var result bool

	if value.Name != "" {
		settings.Services[value.Name] = &value
		result = true
	}

	return result

}

