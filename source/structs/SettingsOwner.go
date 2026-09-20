package structs

import "git-evac/types"
import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "sync"

type SettingsOwner struct {
	mutex      sync.RWMutex
	Name       string                     `json:"name"`
	Identities map[string]*types.Identity `json:"identities"`
	Remotes    map[string]*types.Remote   `json:"remotes"`
	Services   map[string]*types.Service  `json:"services"`
}

func (settings *SettingsOwner) MarshalJSON() ([]byte, error) {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	type Alias SettingsOwner

	return json.Marshal((*Alias)(settings))

}

func (settings *SettingsOwner) IsValid() bool {

	settings.mutex.RLock()

	identities := make(map[string]*types.Identity, len(settings.Identities))
	remotes    := make(map[string]*types.Remote, len(settings.Remotes))

	for name, identity := range settings.Identities {
		identities[name] = identity
	}

	for name, remote := range settings.Remotes {
		remotes[name] = remote
	}

	settings.mutex.RUnlock()

	valid_identities := true
	valid_remotes := true

	for name, identity := range identities {

		if utils_strings.IsName(name) && identity.IsValid() == false {
			valid_identities = false
			break
		}

	}

	for name, remote := range remotes {

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

		settings.mutex.RLock()

		identity, ok := settings.Identities[name]

		if ok == true {
			result = identity
		}

		settings.mutex.RUnlock()

	}

	return result

}

func (settings *SettingsOwner) GetRemote(name string) *types.Remote {

	var result *types.Remote = nil

	if name != "" {

		settings.mutex.RLock()

		remote, ok := settings.Remotes[name]

		if ok == true {
			result = remote
		}

		settings.mutex.RUnlock()

	}

	return result

}

func (settings *SettingsOwner) GetService(name string) *types.Service {

	var result *types.Service = nil

	if name != "" {

		settings.mutex.RLock()

		service, ok := settings.Services[name]

		if ok == true {
			result = service
		}

		settings.mutex.RUnlock()

	}

	return result

}

func (settings *SettingsOwner) RemoveIdentity(name string) bool {

	var result bool

	if name != "" {

		settings.mutex.Lock()

		_, ok := settings.Identities[name]

		if ok == true {
			delete(settings.Identities, name)
			result = true
		}

		settings.mutex.Unlock()

	}

	return result

}

func (settings *SettingsOwner) RemoveRemote(name string) bool {

	var result bool

	if name != "" {

		settings.mutex.Lock()

		_, ok := settings.Remotes[name]

		if ok == true {
			delete(settings.Remotes, name)
			result = true
		}

		settings.mutex.Unlock()

	}

	return result

}

func (settings *SettingsOwner) RemoveService(name string) bool {

	var result bool

	if name != "" {

		settings.mutex.Lock()

		_, ok := settings.Services[name]

		if ok == true {
			delete(settings.Services, name)
			result = true
		}

		settings.mutex.Unlock()

	}

	return result

}

func (settings *SettingsOwner) SetIdentity(value *types.Identity) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Identities[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsOwner) SetRemote(value *types.Remote) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Remotes[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsOwner) SetService(value *types.Service) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Services[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsOwner) SnapshotIdentities() map[string]*types.Identity {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Identity, len(settings.Identities))

	for name, identity := range settings.Identities {
		result[name] = identity
	}

	return result

}

func (settings *SettingsOwner) SnapshotRemotes() map[string]*types.Remote {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Remote, len(settings.Remotes))

	for name, remote := range settings.Remotes {
		result[name] = remote
	}

	return result

}

func (settings *SettingsOwner) SnapshotServices() map[string]*types.Service {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Service, len(settings.Services))

	for name, service := range settings.Services {
		result[name] = service
	}

	return result

}
