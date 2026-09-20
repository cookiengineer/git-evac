package structs

import "git-evac/types"
import "os"
import "strings"

func (profile *Profile) RefreshBackups() {

	backup_folder := profile.Settings.GetBackup()

	stat, err := os.Stat(backup_folder)

	if err == nil && stat.IsDir() {

		profile.Console.Group("Refresh Backups")

		info_owners, err_owners := os.ReadDir(backup_folder)

		if err_owners == nil {

			owners_on_disk := make(map[string]bool)
			backups_on_disk := make(map[string]bool)

			for _, info_owner := range info_owners {

				if info_owner.IsDir() == true {

					owner_name := info_owner.Name()

					owners_on_disk[owner_name] = true

					info_backups, err_backups := os.ReadDir(backup_folder + "/" + info_owner.Name())

					if err_backups == nil {

						for _, info_backup := range info_backups {

							filename := info_backup.Name()

							if info_backup.IsDir() == false && strings.HasSuffix(filename, ".tar.gz") {

								backup_name := filename[0:len(filename)-7]

								backups_on_disk[owner_name + "/" + backup_name] = true

								if profile.HasBackupOwner(owner_name) == false {
									profile.AddBackupOwner(owner_name, backup_folder + "/" + owner_name)
								}

								if profile.HasBackup(owner_name, backup_name) == false {

									owner := profile.GetBackupOwner(owner_name)

									if owner != nil {
										owner.AddBackup(backup_name)
										profile.Console.Log("> " + owner_name + "/" + backup_name)
									}

								}

							}

						}

					}

				}

			}

			// Only remove the in-memory representation. Deleting backups on
			// disk is a manual user task, so no RemoveBackup method exists.
			for owner_name, owner := range profile.SnapshotBackups() {

				if owners_on_disk[owner_name] == false {

					profile.mutex.Lock()
					delete(profile.Backups, owner_name)
					profile.mutex.Unlock()

					profile.Console.Log("> Remove " + owner_name)

				} else {

					for backup_name := range owner.SnapshotBackups() {

						if backups_on_disk[owner_name + "/" + backup_name] == false {

							owner.mutex.Lock()
							delete(owner.Backups, backup_name)
							owner.mutex.Unlock()

							profile.Console.Log("> Remove " + owner_name + "/" + backup_name)

						}

					}

				}

			}

		}

		profile.Console.GroupEnd("Refresh Backups")

	} else {
		profile.Console.Warn("No Backups in Folder \"" + backup_folder + "\"")
	}

	for _, owner := range profile.SnapshotBackups() {

		for _, backup := range owner.SnapshotBackups() {
			backup.Status()
		}

	}

}

func (profile *Profile) AddBackupOwner(owner_name string, owner_folder string) bool {

	var result bool

	profile.mutex.Lock()

	_, ok := profile.Backups[owner_name]

	if ok == false {

		profile.Backups[owner_name] = NewBackupOwner(owner_name, owner_folder)
		result = true

	}

	profile.mutex.Unlock()

	return result

}

func (profile *Profile) GetBackupOwner(owner_name string) *BackupOwner {

	var result *BackupOwner

	profile.mutex.RLock()

	owner, ok := profile.Backups[owner_name]

	if ok == true {
		result = owner
	}

	profile.mutex.RUnlock()

	return result

}

func (profile *Profile) GetBackup(owner_name string, repo_name string) *types.Backup {

	var result *types.Backup

	owner := profile.GetBackupOwner(owner_name)

	if owner != nil {
		result = owner.GetBackup(repo_name)
	}

	return result

}

func (profile *Profile) HasBackupOwner(owner_name string) bool {

	var result bool

	profile.mutex.RLock()

	_, ok := profile.Backups[owner_name]

	if ok == true {
		result = true
	}

	profile.mutex.RUnlock()

	return result

}

func (profile *Profile) HasBackup(owner_name string, repo_name string) bool {

	var result bool

	owner := profile.GetBackupOwner(owner_name)

	if owner != nil {
		result = owner.HasBackup(repo_name)
	}

	return result

}
