package structs

import "io/fs"
import os_user "os/user"

import "fmt"

type Profile struct {
	Backups      map[string]*BackupOwner     `json:"backups"`
	Repositories map[string]*RepositoryOwner `json:"repositories"`
	Settings     *Settings                   `json:"settings"`
	Console      *Console                    `json:"console"`
	Filesystem   *fs.FS                      `json:"-"`
}

func NewProfile(console *Console, settings *Settings) *Profile {

	var profile Profile

	profile.Backups = make(map[string]*BackupOwner)
	profile.Repositories = make(map[string]*RepositoryOwner)

	if console != nil {
		profile.Console = console
	} else {
		profile.Console = NewConsole(nil, nil, 0)
	}

	if settings != nil {

		profile.Settings = settings

	} else {

		user, err := os_user.Current()

		if err == nil {
			profile.Settings = NewSettings(user.HomeDir + "/Backup", user.HomeDir + "/Software", 3000)
		}

	}

	profile.Filesystem = nil

	return &profile

}

func (profile *Profile) Update(settings Settings) {

	// TODO
	fmt.Println(settings)

}

func (profile *Profile) Refresh() {

	profile.RefreshBackups()
	profile.RefreshLocalRepositories()
	profile.RefreshServiceRepositories()

	for _, owner := range profile.Repositories {

		for _, repo := range owner.Repositories {
			repo.Status()
		}

	}

}

