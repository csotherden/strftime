package strftime

import "time"

func Format(format string, t time.Time) string {
	layout := parseFormat(format, convSpecs)
	timeString := t.Format(layout)

	return secondaryFormat(t, timeString)
}
