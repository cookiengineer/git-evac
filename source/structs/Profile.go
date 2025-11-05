package structs

import "io/fs"
import "os"
import "strings"

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

func (profile *Profile) Init() {

	stat0, err0 := os.Stat(profile.Settings.Folder)
	stat1, err1 := os.Stat(profile.Settings.Backup)

	profile.Console.Group("Init()")

	if err0 == nil && stat0.IsDir() {

		profile.Console.Group("Discover Repositories")

		info_owners, err_owners := os.ReadDir(profile.Settings.Folder)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					info_repositories, err_repositories := os.ReadDir(profile.Settings.Folder + "/" + info_owner.Name())

					if err_repositories == nil {

						for _, info_repository := range info_repositories {

							if info_repository.IsDir() == true {

								stat, err := os.Stat(profile.Settings.Folder + "/" + info_owner.Name() + "/" + info_repository.Name() + "/.git")

								if err == nil && stat.IsDir() == true {

									if profile.HasRepositoryOwner(info_owner.Name()) == false {
										profile.AddRepositoryOwner(info_owner.Name(), profile.Settings.Folder + "/" + info_owner.Name())
									}

									if profile.HasRepository(info_owner.Name(), info_repository.Name()) == false {
										owner := profile.GetRepositoryOwner(info_owner.Name())
										owner.AddRepository(info_repository.Name())
										profile.Console.Log("> " + info_owner.Name() + "/" + info_repository.Name())
									}

								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Discover Repositories")

	} else {
		profile.Console.Warn("No Repositories in \"" + profile.Settings.Folder + "\"")
	}

	if err1 == nil && stat1.IsDir() {

		profile.Console.Group("Discover Backups")

		info_owners, err_owners := os.ReadDir(profile.Settings.Backup)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					info_backups, err_backups := os.ReadDir(profile.Settings.Backup + "/" + info_owner.Name())

					if err_backups == nil {

						for _, info_backup := range info_backups {

							if info_backup.IsDir() == false && strings.HasSuffix(info_backup.Name(), ".tar.gz") {

								if profile.HasBackupOwner(info_owner.Name()) == false {
									profile.AddBackupOwner(info_owner.Name(), profile.Settings.Backup + "/" + info_owner.Name())
								}

								if profile.HasBackup(info_owner.Name(), info_backup.Name()) == false {
									owner := profile.GetBackupOwner(info_owner.Name())
									owner.AddBackup(info_backup.Name())
									profile.Console.Log("> " + info_owner.Name() + "/" + info_backup.Name())
								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Discover Backups")

	} else {
		profile.Console.Warn("No Backups in \"" + profile.Settings.Backup + "\"")
	}

	profile.Console.GroupEnd("Init()")

}

func (profile *Profile) Refresh() {

	// TODO: Refresh owners if there are new ones
	// TODO: For each owner refresh repo Status

	for _, owner := range profile.Backups {

		for _, backup := range owner.Backups {
			backup.Status()
		}

	}

	for _, owner := range profile.Repositories {

		for _, repo := range owner.Repositories {
			repo.Status()
		}

	}

}

