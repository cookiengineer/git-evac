package structs

import "git-evac/types"
import "os"
import "strings"

func (profile *Profile) RefreshBackups() {

	stat, err := os.Stat(profile.Settings.Backup)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Backups")

		info_owners, err_owners := os.ReadDir(profile.Settings.Backup)

		if err_owners == nil {

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					info_backups, err_backups := os.ReadDir(profile.Settings.Backup + "/" + info_owner.Name())

					if err_backups == nil {

						for _, info_backup := range info_backups {

							filename := info_backup.Name()

							if info_backup.IsDir() == false && strings.HasSuffix(filename, ".tar.gz") {

								owner_name := info_owner.Name()
								backup_name := filename[0:len(filename)-7]

								if profile.HasBackupOwner(owner_name) == false {
									profile.AddBackupOwner(owner_name, profile.Settings.Backup + "/" + owner_name)
								}

								if profile.HasBackup(owner_name, backup_name) == false {
									owner := profile.GetBackupOwner(owner_name)
									owner.AddBackup(backup_name)
									profile.Console.Log("> " + owner_name + "/" + backup_name)
								}

							}

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Backups")

	} else {
		profile.Console.Warn("No Backups in Folder \"" + profile.Settings.Backup + "\"")
	}

	for _, owner := range profile.Backups {

		for _, backup := range owner.Backups {
			backup.Status()
		}

	}

}

func (profile *Profile) AddBackupOwner(owner_name string, owner_folder string) bool {

	_, ok := profile.Backups[owner_name]

	if ok == false {

		owner := NewBackupOwner(owner_name, owner_folder)
		profile.Backups[owner_name] = &owner

		return true

	}

	return false

}

func (profile *Profile) GetBackupOwner(owner_name string) *BackupOwner {

	var result *BackupOwner

	owner, ok := profile.Backups[owner_name]

	if ok == true {
		result = owner
	}

	return result

}

func (profile *Profile) GetBackup(owner_name string, repo_name string) *types.Backup {

	var result *types.Backup

	owner, ok1 := profile.Backups[owner_name]

	if ok1 == true {

		backup, ok2 := owner.Backups[repo_name]

		if ok2 == true {
			result = backup
		}

	}

	return result

}

func (profile *Profile) HasBackupOwner(owner_name string) bool {

	var result bool

	_, ok := profile.Backups[owner_name]

	if ok == true {
		result = true
	}

	return result

}

func (profile *Profile) HasBackup(owner_name string, repo_name string) bool {

	var result bool

	owner, ok1 := profile.Backups[owner_name]

	if ok1 == true {

		_, ok2 := owner.Backups[repo_name]

		if ok2 == true {
			result = true
		}

	}

	return result

}
