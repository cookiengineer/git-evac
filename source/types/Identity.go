package types

import utils_strings "git-evac/utils/strings"
import "encoding/json"
import "path/filepath"
import "strings"
import "sync"

func toIdentitySSHCommand(sshkey string) string {
	return "ssh -i \"" + sshkey + "\" -F /dev/null"
}

func isValidIdentitySSHKey(sshkey string) bool {

	if sshkey == "" || len(sshkey) > 1024 {
		return false
	}

	if strings.ContainsAny(sshkey, "\n\r\t") {
		return false
	}

	if strings.Contains(sshkey, "\"") || strings.Contains(sshkey, "\\") {
		return false
	}

	if strings.HasPrefix(sshkey, "~/") {

		if strings.Contains(sshkey[2:], "~") {
			return false
		}

	} else if strings.HasPrefix(sshkey, "/") {

		if strings.Contains(sshkey, "~") {
			return false
		}

	} else {
		return false
	}

	if filepath.Clean(sshkey) != sshkey {
		return false
	}

	for s := 0; s < len(sshkey); s++ {

		chr := sshkey[s]

		if chr >= 'a' && chr <= 'z' {
			continue
		} else if chr >= 'A' && chr <= 'Z' {
			continue
		} else if chr >= '0' && chr <= '9' {
			continue
		} else if chr == '/' || chr == '-' || chr == '_' || chr == '.' || chr == '~' {
			continue
		}

		return false

	}

	return true

}

func isValidIdentityGitUserName(username string) bool {

	if username == "" || len(username) > 64 {
		return false
	}

	has_letter := false

	for u := 0; u < len(username); u++ {

		chr := username[u]

		if chr >= 'a' && chr <= 'z' {
			has_letter = true
			continue
		} else if chr >= 'A' && chr <= 'Z' {
			has_letter = true
			continue
		} else if chr >= '0' && chr <= '9' {
			continue
		} else if chr == ' ' || chr == '-' || chr == '_' || chr == '.' || chr == '\'' {
			continue
		}

		return false

	}

	return has_letter

}

type Identity struct {
	mutex  sync.RWMutex
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

func (identity *Identity) MarshalJSON() ([]byte, error) {

	identity.mutex.RLock()
	defer identity.mutex.RUnlock()

	type Alias Identity

	return json.Marshal((*Alias)(identity))

}

func (identity *Identity) GetName() string {

	var result string

	identity.mutex.RLock()
	result = identity.Name
	identity.mutex.RUnlock()

	return result

}

func (identity *Identity) GetSSHKey() string {

	var result string

	identity.mutex.RLock()
	result = identity.SSHKey
	identity.mutex.RUnlock()

	return result

}

func (identity *Identity) GetSSHCommand() string {

	var result string

	identity.mutex.RLock()
	result = identity.Git.Core.SSHCommand
	identity.mutex.RUnlock()

	return result

}

func (identity *Identity) GetGitUserName() string {

	var result string

	identity.mutex.RLock()
	result = identity.Git.User.Name
	identity.mutex.RUnlock()

	return result

}

func (identity *Identity) GetGitUserEmail() string {

	var result string

	identity.mutex.RLock()
	result = identity.Git.User.Email
	identity.mutex.RUnlock()

	return result

}

func (identity *Identity) SetName(value string) bool {

	var result bool

	value = strings.TrimSpace(value)

	if utils_strings.IsName(value) == true {

		identity.mutex.Lock()
		identity.Name = value
		identity.mutex.Unlock()

		result = true

	}

	return result

}

func (identity *Identity) SetSSHKey(value string) bool {

	var result bool

	value = strings.TrimSpace(value)

	if isValidIdentitySSHKey(value) == true {

		identity.mutex.Lock()
		identity.SSHKey = value
		identity.Git.Core.SSHCommand = toIdentitySSHCommand(value)
		identity.mutex.Unlock()

		result = true

	}

	return result

}

func (identity *Identity) SetGitUserName(value string) bool {

	var result bool

	value = strings.TrimSpace(value)

	if isValidIdentityGitUserName(value) == true {

		identity.mutex.Lock()
		identity.Git.User.Name = value
		identity.mutex.Unlock()

		result = true

	}

	return result

}

func (identity *Identity) SetGitUserEmail(value string) bool {

	var result bool

	value = strings.TrimSpace(value)

	if utils_strings.IsEmail(value) == true {

		identity.mutex.Lock()
		identity.Git.User.Email = value
		identity.mutex.Unlock()

		result = true

	}

	return result

}

func (identity *Identity) Sanitize() bool {

	identity.mutex.Lock()
	defer identity.mutex.Unlock()

	name      := strings.TrimSpace(identity.Name)
	sshkey    := strings.TrimSpace(identity.SSHKey)
	username  := strings.TrimSpace(identity.Git.User.Name)
	useremail := strings.TrimSpace(identity.Git.User.Email)

	if utils_strings.IsName(name) == false {
		return false
	}

	if isValidIdentitySSHKey(sshkey) == false {
		return false
	}

	if isValidIdentityGitUserName(username) == false {
		return false
	}

	if utils_strings.IsEmail(useremail) == false {
		return false
	}

	identity.Name = name
	identity.SSHKey = sshkey
	identity.Git.Core.SSHCommand = toIdentitySSHCommand(sshkey)
	identity.Git.User.Name = username
	identity.Git.User.Email = useremail

	return true

}

func (identity *Identity) IsValid() bool {

	// Sanitizing normalizes the untrusted input and derives the SSH command
	return identity.Sanitize()

}
