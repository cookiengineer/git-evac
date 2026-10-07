package structs

import "git-evac/types"
import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "sync"

// SettingsRepository holds per-repository overrides for an owner's template
// settings. Any identity/remote/service defined here wins over the owner-level
// defaults when settings are applied to the repository's git config.
type SettingsRepository struct {
	mutex           sync.RWMutex
	Name            string                     `json:"name"`
	DefaultIdentity string                     `json:"identity"`
	Identities      map[string]*types.Identity `json:"identities"`
	Remotes         map[string]*types.Remote   `json:"remotes"`
	Services        map[string]*types.Service  `json:"services"`
}

func NewSettingsRepository(name string) *SettingsRepository {

	var settings SettingsRepository

	settings.Name = name
	settings.Identities = make(map[string]*types.Identity)
	settings.Remotes = make(map[string]*types.Remote)
	settings.Services = make(map[string]*types.Service)

	return &settings

}

func (settings *SettingsRepository) GetName() string {

	var result string

	settings.mutex.RLock()
	result = settings.Name
	settings.mutex.RUnlock()

	return result

}

func (settings *SettingsRepository) GetDefaultIdentity() string {

	var result string

	settings.mutex.RLock()
	result = settings.DefaultIdentity
	settings.mutex.RUnlock()

	return result

}

func (settings *SettingsRepository) SetDefaultIdentity(value string) {

	settings.mutex.Lock()
	settings.DefaultIdentity = value
	settings.mutex.Unlock()

}

func (settings *SettingsRepository) MarshalJSON() ([]byte, error) {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	type Alias SettingsRepository

	return json.Marshal((*Alias)(settings))

}

func (settings *SettingsRepository) UnmarshalJSON(data []byte) error {

	settings.mutex.Lock()
	defer settings.mutex.Unlock()

	type Alias SettingsRepository
	alias := (*Alias)(settings)

	err := json.Unmarshal(data, alias)

	if err != nil {
		return err
	}

	if settings.Identities == nil {
		settings.Identities = make(map[string]*types.Identity)
	}

	if settings.Remotes == nil {
		settings.Remotes = make(map[string]*types.Remote)
	}

	if settings.Services == nil {
		settings.Services = make(map[string]*types.Service)
	}

	return nil

}

func (settings *SettingsRepository) IsValid() bool {

	settings.mutex.RLock()

	identities := make(map[string]*types.Identity, len(settings.Identities))
	remotes := make(map[string]*types.Remote, len(settings.Remotes))

	for name, identity := range settings.Identities {
		identities[name] = identity
	}

	for name, remote := range settings.Remotes {
		remotes[name] = remote
	}

	settings.mutex.RUnlock()

	valid_identities := true
	valid_remotes := true

	// A nil entry in a pinned map is a removal marker and therefore valid.
	for name, identity := range identities {

		if utils_strings.IsName(name) && identity != nil && identity.IsValid() == false {
			valid_identities = false
			break
		}

	}

	for name, remote := range remotes {

		if utils_strings.IsName(name) && remote != nil && remote.IsValid() == false {
			valid_remotes = false
			break
		}

	}

	return valid_identities && valid_remotes

}

func (settings *SettingsRepository) GetIdentity(name string) *types.Identity {

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

func (settings *SettingsRepository) GetRemote(name string) *types.Remote {

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

func (settings *SettingsRepository) GetService(name string) *types.Service {

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

func (settings *SettingsRepository) RemoveIdentity(name string) bool {

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

func (settings *SettingsRepository) RemoveRemote(name string) bool {

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

func (settings *SettingsRepository) RemoveService(name string) bool {

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

func (settings *SettingsRepository) SetIdentity(value *types.Identity) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Identities[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsRepository) SetRemote(value *types.Remote) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Remotes[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsRepository) SetService(value *types.Service) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Services[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsRepository) SnapshotIdentities() map[string]*types.Identity {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Identity, len(settings.Identities))

	for name, identity := range settings.Identities {
		result[name] = identity
	}

	return result

}

func (settings *SettingsRepository) SnapshotRemotes() map[string]*types.Remote {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Remote, len(settings.Remotes))

	for name, remote := range settings.Remotes {
		result[name] = remote
	}

	return result

}

func (settings *SettingsRepository) SnapshotServices() map[string]*types.Service {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*types.Service, len(settings.Services))

	for name, service := range settings.Services {
		result[name] = service
	}

	return result

}
