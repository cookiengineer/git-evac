package structs

import "git-evac/types"
import "encoding/json"
import "sync"

type BackupOwner struct {
	mutex   sync.RWMutex
	Name    string                   `json:"name"`
	Folder  string                   `json:"folder"`
	Backups map[string]*types.Backup `json:"backups"`
}

func NewBackupOwner(name string, folder string) *BackupOwner {

	var owner BackupOwner

	owner.Name = name
	owner.Folder = folder
	owner.Backups = make(map[string]*types.Backup)

	return &owner

}

func (owner *BackupOwner) MarshalJSON() ([]byte, error) {

	owner.mutex.RLock()
	defer owner.mutex.RUnlock()

	type Alias BackupOwner

	return json.Marshal((*Alias)(owner))

}

func (owner *BackupOwner) AddBackup(name string) bool {

	var result bool

	owner.mutex.Lock()

	_, ok := owner.Backups[name]

	if ok == false {

		owner.Backups[name] = types.NewBackup(name, owner.Folder + "/" + name + ".tar.gz")
		result = true

	}

	owner.mutex.Unlock()

	return result

}

func (owner *BackupOwner) GetBackup(name string) *types.Backup {

	var result *types.Backup

	owner.mutex.RLock()

	tmp, ok := owner.Backups[name]

	if ok == true {
		result = tmp
	}

	owner.mutex.RUnlock()

	return result

}

func (owner *BackupOwner) HasBackup(name string) bool {

	var result bool

	owner.mutex.RLock()

	_, ok := owner.Backups[name]

	if ok == true {
		result = true
	}

	owner.mutex.RUnlock()

	return result

}

func (owner *BackupOwner) SnapshotBackups() map[string]*types.Backup {

	owner.mutex.RLock()
	defer owner.mutex.RUnlock()

	result := make(map[string]*types.Backup, len(owner.Backups))

	for name, backup := range owner.Backups {
		result[name] = backup
	}

	return result

}
