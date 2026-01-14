package structs

import utils_paths "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"

type Settings struct {
	Backup string                   `json:"backup"`
	Folder string                   `json:"folder"`
	Port   uint16                   `json:"port"`
	Owners map[string]SettingsOwner `json:"owners"`
}

func NewSettings(backup string, folder string, port uint16) *Settings {

	var settings Settings

	settings.Backup = backup
	settings.Folder = folder
	settings.Port   = port
	settings.Owners = make(map[string]SettingsOwner)

	return &settings

}

func (settings *Settings) IsValid() bool {

	if settings.Backup != "" && settings.Folder != "" && settings.Port != 0 {

		valid_backup := false
		valid_folder := false
		valid_port := false
		valid_owners := true

		if utils_paths.IsFolder(settings.Backup) {
			valid_backup = true
		}

		if utils_paths.IsFolder(settings.Folder) {
			valid_folder = true
		}

		if settings.Port > 1025 && settings.Port < 65535 {
			valid_port = true
		}

		for name, owner := range settings.Owners {

			if utils_strings.IsName(name) && owner.IsValid() == false {
				valid_owners = false
				break
			}

		}

		return valid_backup && valid_folder && valid_port && valid_owners

	}

	return false

}
