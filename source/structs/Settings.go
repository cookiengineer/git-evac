package structs

import utils_paths "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"

type Settings struct {
	Backup        string                          `json:"backup"`
	Folder        string                          `json:"folder"`
	Port          uint16                          `json:"port"`
	Organizations map[string]OrganizationSettings `json:"organizations"`
}

func (settings *Settings) IsValid() bool {

	if settings.Backup != "" && settings.Folder != "" && settings.Port != 0 {

		valid_backup := false
		valid_folder := false
		valid_port := false
		valid_organizations := true

		if utils_paths.IsFolder(settings.Backup) {
			valid_backup = true
		}

		if utils_paths.IsFolder(settings.Folder) {
			valid_folder = true
		}

		if settings.Port > 1025 && settings.Port < 65535 {
			valid_port = true
		}

		for name, orga := range settings.Organizations {

			if utils_strings.IsName(name) && orga.IsValid() == false {
				valid_organizations = false
				break
			}

		}

		return valid_backup && valid_folder && valid_port && valid_organizations

	}

	return false

}
