package components

import "fmt"
import "time"

func formatDuration(start time.Time, stop time.Time) string {

	if !start.IsZero() && !stop.IsZero() {

		duration := stop.Sub(start)
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)

	} else if !start.IsZero() && stop.IsZero() {

		duration := time.Now().Sub(start)
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60

		return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)

	} else {
		return "00h 00m 00s"
	}

}
