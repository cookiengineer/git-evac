package structs

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

func (profile *Profile) GetBackup(owner_name string, repo_name string) *Backup {

	var result *Backup

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
