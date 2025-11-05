package schemas

import "git-evac/structs"

type Backups struct {
	Owners map[string]*structs.BackupOwner `json:"owners"`
}
