package strftime

import "time"

func Parse(format, timeString string) (time.Time, error) {
	layout := parseFormat(format, convSpecs)

	return time.Parse(layout, timeString)
}
