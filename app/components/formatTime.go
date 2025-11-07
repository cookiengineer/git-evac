package components

import "time"

func formatTime(datetime time.Time) string {
	return datetime.Format(time.DateTime)
}
