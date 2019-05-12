package strftime

import "time"

func Convert(f, t string) (time.Time, error) {
	layout := parseFormat(f, convSpecs)

	return time.Parse(layout, t)
}
