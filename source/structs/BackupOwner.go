package structs

import "git-evac/types"

type BackupOwner struct {
	Name    string                   `json:"name"`
	Folder  string                   `json:"folder"`
	Backups map[string]*types.Backup `json:"backups"`
}

func NewBackupOwner(name string, folder string) BackupOwner {

	var owner BackupOwner

	owner.Name = name
	owner.Folder = folder
	owner.Backups = make(map[string]*types.Backup)

	return owner

}

func (owner *BackupOwner) AddBackup(name string) bool {

	_, ok := owner.Backups[name]

	if ok == false {

		owner.Backups[name] = types.NewBackup(name, owner.Folder + "/" + name + ".tar.gz")

		return true

	}

	return false

}

func (owner *BackupOwner) GetBackup(name string) *types.Backup {

	var result *types.Backup

	tmp, ok := owner.Backups[name]

	if ok == true {
		result = tmp
	}

	return result

}

func (owner *BackupOwner) HasBackup(name string) bool {

	var result bool

	_, ok := owner.Backups[name]

	if ok == true {
		result = true
	}

	return result

}

