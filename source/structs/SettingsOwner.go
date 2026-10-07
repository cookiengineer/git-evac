package structs

import "git-evac/types"
import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "sync"

type SettingsOwner struct {
	mutex           sync.RWMutex
	Name            string                         `json:"name"`
	DefaultIdentity string                         `json:"identity"`
	Identities      map[string]*types.Identity     `json:"identities"`
	Remotes         map[string]*types.Remote       `json:"remotes"`
	Services        map[string]*types.Service      `json:"services"`
	Repositories    map[string]*SettingsRepository `json:"repositories"`
}

func NewSettingsOwner(name string) *SettingsOwner {

	var settings SettingsOwner

	settings.Name = name
	settings.DefaultIdentity = ""
	settings.Identities = make(map[string]*types.Identity)
	settings.Remotes = make(map[string]*types.Remote)
	settings.Services = make(map[string]*types.Service)
	settings.Repositories = make(map[string]*SettingsRepository)

	return &settings

}

func (settings *SettingsOwner) MarshalJSON() ([]byte, error) {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	type Alias SettingsOwner

	return json.Marshal((*Alias)(settings))

}

func (settings *SettingsOwner) UnmarshalJSON(data []byte) error {

	settings.mutex.Lock()
	defer settings.mutex.Unlock()

	type Alias SettingsOwner
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

	if settings.Repositories == nil {
		settings.Repositories = make(map[string]*SettingsRepository)
	}

	return nil

}

func (settings *SettingsOwner) IsValid() bool {

	settings.mutex.RLock()

	identities := make(map[string]*types.Identity, len(settings.Identities))
	remotes := make(map[string]*types.Remote, len(settings.Remotes))
	repositories := make(map[string]*SettingsRepository, len(settings.Repositories))

	for name, identity := range settings.Identities {
		identities[name] = identity
	}

	for name, remote := range settings.Remotes {
		remotes[name] = remote
	}

	for name, repository := range settings.Repositories {
		repositories[name] = repository
	}

	settings.mutex.RUnlock()

	valid_identities := true
	valid_remotes := true
	valid_repositories := true

	for name, identity := range identities {

		if utils_strings.IsName(name) && (identity == nil || identity.IsValid() == false) {
			valid_identities = false
			break
		}

	}

	for name, remote := range remotes {

		if utils_strings.IsName(name) && (remote == nil || remote.IsValid() == false) {
			valid_remotes = false
			break
		}

	}

	for name, repository := range repositories {

		if utils_strings.IsName(name) && (repository == nil || repository.IsValid() == false) {
			valid_repositories = false
			break
		}

	}

	return valid_identities && valid_remotes && valid_repositories

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

func (settings *SettingsOwner) GetRepository(name string) *SettingsRepository {

	var result *SettingsRepository = nil

	if name != "" {

		settings.mutex.RLock()

		repository, ok := settings.Repositories[name]

		if ok == true {
			result = repository
		}

		settings.mutex.RUnlock()

	}

	return result

}

func (settings *SettingsOwner) SetRepository(value *SettingsRepository) bool {

	var result bool

	if value != nil && value.GetName() != "" {

		settings.mutex.Lock()
		settings.Repositories[value.GetName()] = value
		settings.mutex.Unlock()

		result = true

	}

	return result

}

func (settings *SettingsOwner) SetRepositories(value map[string]*SettingsRepository) {

	settings.mutex.Lock()
	settings.Repositories = value
	settings.mutex.Unlock()

}

func (settings *SettingsOwner) RemoveRepository(name string) bool {

	var result bool

	if name != "" {

		settings.mutex.Lock()

		_, ok := settings.Repositories[name]

		if ok == true {
			delete(settings.Repositories, name)
			result = true
		}

		settings.mutex.Unlock()

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

func (settings *SettingsOwner) SnapshotRepositories() map[string]*SettingsRepository {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*SettingsRepository, len(settings.Repositories))

	for name, repository := range settings.Repositories {
		result[name] = repository
	}

	return result

}

// Effective merges the owner-level template settings with the per-repository
// pinned overrides and returns the settings that should be applied to the git
// config of the given repository. A pinned entry with a nil value removes the
// inherited owner-level entry of the same name.
func (settings *SettingsOwner) Effective(repository_name string) *SettingsRepository {

	var result *SettingsRepository = nil

	if repository_name != "" {

		settings.mutex.RLock()

		owner_identities := make(map[string]*types.Identity, len(settings.Identities))
		owner_remotes := make(map[string]*types.Remote, len(settings.Remotes))
		owner_services := make(map[string]*types.Service, len(settings.Services))
		owner_identity := settings.DefaultIdentity

		for name, identity := range settings.Identities {
			owner_identities[name] = identity
		}

		for name, remote := range settings.Remotes {
			owner_remotes[name] = remote
		}

		for name, service := range settings.Services {
			owner_services[name] = service
		}

		pinned, _ := settings.Repositories[repository_name]

		settings.mutex.RUnlock()

		result = NewSettingsRepository(repository_name)

		for name, identity := range owner_identities {
			result.Identities[name] = identity
		}

		for name, remote := range owner_remotes {
			result.Remotes[name] = remote
		}

		for name, service := range owner_services {
			result.Services[name] = service
		}

		result.DefaultIdentity = owner_identity

		if pinned != nil {

			for name, identity := range pinned.SnapshotIdentities() {

				if identity == nil {
					delete(result.Identities, name)
				} else {
					result.Identities[name] = identity
				}

			}

			for name, remote := range pinned.SnapshotRemotes() {

				if remote == nil {
					delete(result.Remotes, name)
				} else {
					result.Remotes[name] = remote
				}

			}

			for name, service := range pinned.SnapshotServices() {

				if service == nil {
					delete(result.Services, name)
				} else {
					result.Services[name] = service
				}

			}

			if pinned.GetDefaultIdentity() != "" {
				result.DefaultIdentity = pinned.GetDefaultIdentity()
			}

		}

	}

	return result

}

// EffectiveIdentity returns the identity that should be applied to the given
// repository, honoring a pinned identity override.
func (settings *SettingsOwner) EffectiveIdentity(repository_name string) *types.Identity {

	var result *types.Identity = nil

	effective := settings.Effective(repository_name)

	if effective != nil && effective.GetDefaultIdentity() != "" {
		result = effective.GetIdentity(effective.GetDefaultIdentity())
	}

	return result

}
