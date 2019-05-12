package strftime

import "time"

// Format returns the provided time.Time formatted according to the strftime(3) based format string
func Format(format string, t time.Time) string {
	layout := parseFormat(format, convSpecs)
	timeString := t.Format(layout)

	return secondaryFormat(t, timeString)
}
