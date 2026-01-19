package types

import "os"
import "strings"
import "time"

type Backup struct {
	Name string    `json:"name"`
	File string    `json:"file"` // /path/to/file.tar.gz
	Size int64     `json:"size"`
	Time time.Time `json:"time"`
}

func NewBackup(name string, file string) *Backup {

	var backup Backup

	backup.Name = name
	backup.File = file
	backup.Size = 0
	backup.Time = time.Time{}

	return &backup

}

func (backup *Backup) Status() bool {

	stat, err0 := os.Stat(backup.File)

	if err0 == nil && stat.Mode().IsRegular() && strings.HasSuffix(backup.File, ".tar.gz") {

		backup.Size = stat.Size()
		backup.Time = stat.ModTime()

		return true

	}

	return false

}
