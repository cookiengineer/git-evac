package structs

import utils_paths "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "sync"

type Settings struct {
	mutex  sync.RWMutex
	Backup string                    `json:"backup"`
	Folder string                    `json:"folder"`
	Port   uint16                    `json:"port"`
	Owners map[string]*SettingsOwner `json:"owners"`
}

func NewSettings(backup string, folder string, port uint16) *Settings {

	var settings Settings

	settings.Backup = backup
	settings.Folder = folder
	settings.Port   = port
	settings.Owners = make(map[string]*SettingsOwner)

	return &settings

}

func (settings *Settings) MarshalJSON() ([]byte, error) {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	type Alias Settings

	return json.Marshal((*Alias)(settings))

}

func (settings *Settings) GetBackup() string {

	var result string

	settings.mutex.RLock()
	result = settings.Backup
	settings.mutex.RUnlock()

	return result

}

func (settings *Settings) GetFolder() string {

	var result string

	settings.mutex.RLock()
	result = settings.Folder
	settings.mutex.RUnlock()

	return result

}

func (settings *Settings) GetPort() uint16 {

	var result uint16

	settings.mutex.RLock()
	result = settings.Port
	settings.mutex.RUnlock()

	return result

}

func (settings *Settings) GetOwner(name string) *SettingsOwner {

	var result *SettingsOwner = nil

	if name != "" {

		settings.mutex.RLock()

		owner, ok := settings.Owners[name]

		if ok == true {
			result = owner
		}

		settings.mutex.RUnlock()

	}

	return result

}

func (settings *Settings) GetOwners() map[string]*SettingsOwner {

	settings.mutex.RLock()
	defer settings.mutex.RUnlock()

	result := make(map[string]*SettingsOwner, len(settings.Owners))

	for name, owner := range settings.Owners {
		result[name] = owner
	}

	return result

}

func (settings *Settings) SetBackup(value string) {

	settings.mutex.Lock()
	settings.Backup = value
	settings.mutex.Unlock()

}

func (settings *Settings) SetFolder(value string) {

	settings.mutex.Lock()
	settings.Folder = value
	settings.mutex.Unlock()

}

func (settings *Settings) SetPort(value uint16) {

	settings.mutex.Lock()
	settings.Port = value
	settings.mutex.Unlock()

}

func (settings *Settings) SetOwners(value map[string]*SettingsOwner) {

	settings.mutex.Lock()
	settings.Owners = value
	settings.mutex.Unlock()

}

func (settings *Settings) IsValid() bool {

	settings.mutex.RLock()

	backup := settings.Backup
	folder := settings.Folder
	port   := settings.Port

	owners := make(map[string]*SettingsOwner, len(settings.Owners))

	for name, owner := range settings.Owners {
		owners[name] = owner
	}

	settings.mutex.RUnlock()

	if backup != "" && folder != "" && port != 0 {

		valid_backup := false
		valid_folder := false
		valid_port := false
		valid_owners := true

		if utils_paths.IsFolder(backup) {
			valid_backup = true
		}

		if utils_paths.IsFolder(folder) {
			valid_folder = true
		}

		if port > 1025 && port < 65535 {
			valid_port = true
		}

		for name, owner := range owners {

			if utils_strings.IsName(name) && owner.IsValid() == false {
				valid_owners = false
				break
			}

		}

		return valid_backup && valid_folder && valid_port && valid_owners

	}

	return false

}
