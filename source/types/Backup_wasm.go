//go:build wasm

package types

import "strings"

func (backup *Backup) Status() bool {

	if strings.HasSuffix(backup.File, ".tar.gz") {
		return true
	}

	return false

}
