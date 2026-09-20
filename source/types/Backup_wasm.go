//go:build wasm

package types

import "strings"

func (backup *Backup) Status() bool {

	backup.mutex.Lock()
	defer backup.mutex.Unlock()

	if strings.HasSuffix(backup.File, ".tar.gz") {
		return true
	}

	return false

}
