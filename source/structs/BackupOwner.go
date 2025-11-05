package structs

type BackupOwner struct {
	Name    string             `json:"name"`
	Folder  string             `json:"folder"`
	Backups map[string]*Backup `json:"backups"`
}

func NewBackupOwner(name string, folder string) BackupOwner {

	var owner BackupOwner

	owner.Name = name
	owner.Folder = folder
	owner.Backups = make(map[string]*Backup)

	return owner

}

func (owner *BackupOwner) AddBackup(name string) bool {

	_, ok := owner.Backups[name]

	if ok == false {

		repo := NewBackup(name, owner.Folder + "/" + name + ".tar.gz")
		owner.Backups[name] = &repo

		return true

	}

	return false

}

func (owner *BackupOwner) GetBackup(name string) *Backup {

	var result *Backup

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

