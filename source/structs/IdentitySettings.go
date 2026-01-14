package structs

import utils_paths   "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"
import "path/filepath"
import "strings"

type IdentitySettings struct {
	Name   string `json:"name"`
	SSHKey string `json:"ssh-key"`
	Git struct {
		Core struct {
			// git config --file .git/config core.sshCommand "ssh -i \"/home/cookiengineer/.ssh/identity.key\" -F /dev/null"
			SSHCommand string `json:"sshCommand"`
		} `json:"core"`
		User struct {
			// git config --file .git/config user.name  "John Doe"
			Name  string `json:"name"`
			// git config --file .git/config user.email john@example.com
			Email string `json:"email"`
		} `json:"user"`
	} `json:"git"`

}

func NewIdentitySettings(name string) IdentitySettings {

	var settings IdentitySettings

	settings.Name = name
	settings.SSHKey = ""
	settings.Git.Core.SSHCommand = ""
	settings.Git.User.Name = ""
	settings.Git.User.Email = ""

	return settings

}

func (settings *IdentitySettings) IsValid() bool {

	var result bool

	if utils_strings.IsName(settings.Name) {

		valid_sshkey := false
		valid_git_core := false
		valid_git_user := false

		folder := filepath.Dir(settings.SSHKey)

		if utils_paths.IsFolder(folder) {
			valid_sshkey = true
		}

		if strings.HasPrefix(settings.Git.Core.SSHCommand, "ssh -i \"") && strings.HasSuffix(settings.Git.Core.SSHCommand, "\" -F /dev/null") {

			sshkey_file := settings.Git.Core.SSHCommand[8:len(settings.Git.Core.SSHCommand)-14]

			if strings.HasPrefix(sshkey_file, "/") && sshkey_file == settings.SSHKey {
				valid_git_core = true
			}

		}

		if strings.Contains(settings.Git.User.Name, " ") {

			if utils_strings.IsEmail(settings.Git.User.Email) {

				valid_git_user = true

				tmp := strings.Split(settings.Git.User.Name, " ")

				for t := 0; t < len(tmp); t++ {

					if !utils_strings.IsName(strings.ToLower(tmp[t])) {
						valid_git_user = false
						break
					}

				}

			}

		} else if utils_strings.IsName(settings.Git.User.Name) {

			if utils_strings.IsEmail(settings.Git.User.Email) {
				valid_git_user = true
			}

		}

		return valid_sshkey && valid_git_core && valid_git_user

	}

	// TODO: Validate name
	// TODO: Validate key path (being absolute or with ~/ prefix)
	// TODO: Set Git.Core.SSHCommand value
	// TODO: Set Git.User.Name value
	// TODO: Set Git.User.Email value

	return result

}
