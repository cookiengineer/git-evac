//go:build !wasm

package types

import "os"
import "strings"

func (backup *Backup) Status() bool {

	stat, err0 := os.Stat(backup.File)

	if err0 == nil && stat.Mode().IsRegular() && strings.HasSuffix(backup.File, ".tar.gz") {

		backup.Size = stat.Size()
		backup.Time = stat.ModTime()

		return true

	}

	return false

}
