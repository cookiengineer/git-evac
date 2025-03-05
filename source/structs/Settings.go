package structs

import utils_paths "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"

type Settings struct {
	Backup     string                      `json:"backup"`
	Folder     string                      `json:"folder"`
	Port       uint16                      `json:"port"`
	Identities map[string]IdentitySettings `json:"identities"`
	Remotes    map[string]RemoteSettings   `json:"remotes"`
}

func (settings *Settings) IsValid() bool {

	if settings.Backup != "" && settings.Folder != "" && settings.Port != 0 {

		valid_backup := false
		valid_folder := false
		valid_port := false
		valid_identities := true
		valid_remotes := true

		if utils_paths.IsFolder(settings.Backup) {
			valid_backup = true
		}

		if utils_paths.IsFolder(settings.Folder) {
			valid_folder = true
		}

		if settings.Port > 1025 && settings.Port < 65535 {
			valid_port = true
		}

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

		return valid_backup && valid_folder && valid_port && valid_identities && valid_remotes

	}

	return false

}
