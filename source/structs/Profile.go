package structs

import "io/fs"

type Profile struct {
	Backups      map[string]*BackupOwner     `json:"backups"`
	Repositories map[string]*RepositoryOwner `json:"repositories"`
	Settings     Settings                    `json:"settings"`
	Console      *Console                    `json:"console"`
	Filesystem   *fs.FS                      `json:"-"`
}

func NewProfile(console *Console, backup string, folder string, port uint16) *Profile {

	var profile Profile

	profile.Backups = make(map[string]*BackupOwner)
	profile.Repositories = make(map[string]*RepositoryOwner)

	if console != nil {
		profile.Console = console
	} else {
		profile.Console = NewConsole(nil, nil, 0)
	}

	profile.Filesystem = nil

	profile.Settings.Backup        = backup
	profile.Settings.Folder        = folder
	profile.Settings.Port          = port
	profile.Settings.Organizations = make(map[string]OrganizationSettings)

	return &profile

}

func (profile *Profile) Refresh() {
	profile.RefreshBackups()
	profile.RefreshRepositories()
}

