package types

import utils_paths   "git-evac/utils/paths"
import utils_strings "git-evac/utils/strings"
import "path/filepath"
import "strings"

type Identity struct {
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

func NewIdentity(name string) *Identity {

	var identity Identity

	identity.Name = name
	identity.SSHKey = ""
	identity.Git.Core.SSHCommand = ""
	identity.Git.User.Name = ""
	identity.Git.User.Email = ""

	return &identity

}

func (identity *Identity) IsValid() bool {

	var result bool

	if utils_strings.IsName(identity.Name) {

		valid_sshkey := false
		valid_git_core := false
		valid_git_user := false

		folder := filepath.Dir(identity.SSHKey)

		if utils_paths.IsFolder(folder) {
			valid_sshkey = true
		}

		if strings.HasPrefix(identity.Git.Core.SSHCommand, "ssh -i \"") && strings.HasSuffix(identity.Git.Core.SSHCommand, "\" -F /dev/null") {

			sshkey_file := identity.Git.Core.SSHCommand[8:len(identity.Git.Core.SSHCommand)-14]

			if strings.HasPrefix(sshkey_file, "/") && sshkey_file == identity.SSHKey {
				valid_git_core = true
			}

		}

		if strings.Contains(identity.Git.User.Name, " ") {

			if utils_strings.IsEmail(identity.Git.User.Email) {

				valid_git_user = true

				tmp := strings.Split(identity.Git.User.Name, " ")

				for t := 0; t < len(tmp); t++ {

					if !utils_strings.IsName(strings.ToLower(tmp[t])) {
						valid_git_user = false
						break
					}

				}

			}

		} else if utils_strings.IsName(identity.Git.User.Name) {

			if utils_strings.IsEmail(identity.Git.User.Email) {
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
