package strftime

import "time"

// Parse parses a formatted string and returns the time.Time value it represents.
// The format defines the input value format using C strftime(3) conversion specifications.
func Parse(format, value string) (time.Time, error) {
	layout := parseFormat(format, convSpecs)

	return time.Parse(layout, value)
}
