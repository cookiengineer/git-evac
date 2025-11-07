package components

import "fmt"
import "math"

func formatSize(bytes int64) string {

    if bytes < 1024 {
        return fmt.Sprintf("%4d B", bytes)
    }

	unit := float64(1024.0)
    sizes := []string{"B", "KB", "MB", "GB", "TB", "PB"}

    index := math.Floor(math.Log(float64(bytes)) / math.Log(unit))
    value := float64(bytes) / math.Pow(unit, index)

    return fmt.Sprintf("%7.2f %s", value, sizes[int(index)])

}
