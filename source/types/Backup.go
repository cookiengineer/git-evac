package types

import "encoding/json"
import "sync"
import "time"

type Backup struct {
	mutex sync.RWMutex
	Name  string    `json:"name"`
	File  string    `json:"file"` // /path/to/file.tar.gz
	Size  int64     `json:"size"`
	Time  time.Time `json:"time"`
}

func NewBackup(name string, file string) *Backup {

	var backup Backup

	backup.Name = name
	backup.File = file
	backup.Size = 0
	backup.Time = time.Time{}

	return &backup

}

func (backup *Backup) MarshalJSON() ([]byte, error) {

	backup.mutex.RLock()
	defer backup.mutex.RUnlock()

	type Alias Backup

	return json.Marshal((*Alias)(backup))

}
